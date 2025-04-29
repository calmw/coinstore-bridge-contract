package ethereum

import (
	"coinstore/bridge/chains"
	"coinstore/bridge/chains/signature"
	"coinstore/bridge/event"
	"coinstore/bridge/msg"
	"coinstore/model"
	"coinstore/utils"
	"context"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"os"
	"time"
)

// ExecuteBlockWatchLimit Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 100

// TxRetryInterval Time between retrying a failed tx
const TxRetryInterval = time.Second * 2

// TxRetryLimit Maximum number of tx retries before exiting
const TxRetryLimit = 10

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var ErrFatalTx = errors.New("submission of transaction failed")
var ErrFatalQuery = errors.New("query of chain state failed")

func (w *Writer) proposalIsComplete(m msg.Message, dataHash [32]byte) bool {
	prop, err := w.voteContract.GetProposal(nil, m.Source.Big(), m.DepositNonce.Big(), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	///
	w.log.Debug("🫥  proposal status", "status", prop.Status, "nonce", m.DepositNonce)
	if prop.Status >= 2 {
		model.UpdateVoteStatus(m, 1)
	}
	///
	return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
}

func (w *Writer) proposalIsFinalized(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.voteContract.GetProposal(w.conn.CallOpts(), srcId.Big(), nonce.Big(), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
}

func (w *Writer) proposalIsPassed(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.voteContract.GetProposal(w.conn.CallOpts(), srcId.Big(), nonce.Big(), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus
}

func IDAndNonce(srcId msg.ChainId, nonce msg.Nonce) *big.Int {
	var data []byte
	data = append(data, nonce.Big().Bytes()...)
	data = append(data, uint8(srcId))
	return big.NewInt(0).SetBytes(data)
}

func (w *Writer) hasVoted(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	fmt.Println(w.conn.CallOpts())
	sigAccount := os.Getenv("SIG_ACCOUNT_EVM")
	if len(sigAccount) <= 0 {
		sigAccount = "0x1933ccd14cafe561d862e5f35d0de75322a55412"
	}
	fromAddress := common.HexToAddress(sigAccount)
	hasVoted, err := w.voteContract.HasVotedOnProposal(w.conn.CallOpts(), IDAndNonce(srcId, nonce), dataHash, fromAddress)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}

	return hasVoted
}

func (w *Writer) shouldVote(m msg.Message, dataHash [32]byte) bool {
	if w.proposalIsComplete(m, dataHash) {
		w.log.Info("Proposal complete, not voting", "src", m.Source, "nonce", m.DepositNonce)
		return false
	}

	if w.hasVoted(m.Source, m.DepositNonce, dataHash) {
		w.log.Info("Relayer has already voted, not voting", "src", m.Source, "nonce", m.DepositNonce)
		return false
	}

	return true
}

func (w *Writer) CreateProposal(m msg.Message) bool {
	fmt.Println("--")
	fmt.Println(w.log)
	fmt.Println("--")
	w.log.Info("Creating generic proposal", "src", m.Source, "nonce", m.DepositNonce)

	metadata := m.Payload[0].([]byte)
	data := chains.ConstructGenericProposalData(metadata)
	toHash := append(common.HexToAddress(w.Cfg.BridgeContractAddress).Bytes(), data...)
	dataHash := utils.Keccak256(toHash)
	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			w.ExecuteProposal(m, data, dataHash)
			return true
		} else {
			return false
		}
	}

	latestBlock, err := w.conn.LatestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	go w.watchThenExecute(m, data, dataHash, latestBlock)

	w.voteProposal(m, dataHash)

	return true
}

func (w *Writer) watchThenExecute(m msg.Message, data []byte, dataHash [32]byte, latestBlock *big.Int) {
	w.log.Info("Watching for finalization event", "src", m.Source, "nonce", m.DepositNonce)

	for i := 0; i < ExecuteBlockWatchLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			for waitRetrys := 0; waitRetrys < BlockRetryLimit; waitRetrys++ {
				err := w.conn.WaitForBlock(latestBlock, w.Cfg.BlockConfirmations)
				if err != nil {
					w.log.Error("Waiting for block failed", "err", err)
					if waitRetrys+1 == BlockRetryLimit {
						w.log.Error("Waiting for block retries exceeded, shutting down")
						w.sysErr <- ErrFatalQuery
						return
					}
				} else {
					break
				}
			}

			query := buildQuery(common.HexToAddress(w.Cfg.VoteContractAddress), event.ProposalEvent, latestBlock, latestBlock)
			evts, err := w.conn.ClientEvm().FilterLogs(context.Background(), query)
			if err != nil {
				w.log.Error("Failed to fetch logs", "err", err)
				return
			}

			for _, evt := range evts {
				sourceId := evt.Topics[1].Big().Uint64()
				depositNonce := evt.Topics[2].Big().Uint64()
				status := evt.Topics[3].Big().Uint64()

				if m.Source == msg.ChainId(sourceId) &&
					m.DepositNonce.Big().Uint64() == depositNonce &&
					event.IsFinalized(uint8(status)) {
					w.ExecuteProposal(m, data, dataHash)
					return
				} else {
					w.log.Trace("Ignoring event", "src", sourceId, "nonce", depositNonce)
				}
			}
			w.log.Trace("No finalization event found in current block", "block", latestBlock, "src", m.Source, "nonce", m.DepositNonce)
			latestBlock = latestBlock.Add(latestBlock, big.NewInt(1))
		}
	}
	log.Warn("Block watch limit exceeded, skipping execution", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
}

//func (w *Writer) voteProposal(m msg.Message, dataHash [32]byte) {
//	w.muVote.Lock()
//	defer w.muVote.Unlock()
//
//	for i := 0; i < TxRetryLimit; i++ {
//		select {
//		case <-w.stop:
//			return
//		default:
//			vStatus, _, err := model.GetBridgeTxStatus(m)
//			if err == nil && vStatus > 0 {
//				w.log.Info("voteProposal", "skip, src", m.Source, "depositNonce", m.DepositNonce)
//				return
//			}
//			//err = w.conn.LockAndUpdateOpts()
//			//if err != nil {
//			//	w.log.Error("Failed to update tx opts", "err", err)
//			//	continue
//			//}
//			//
//			//gasLimit := w.conn.Opts().GasLimit
//			//gasPrice := w.conn.Opts().GasPrice
//			//
//			//w.log.Debug("voteProposal", "dataHash", fmt.Sprintf("%x", dataHash))
//			//w.log.Debug("voteProposal", "DepositNonce", fmt.Sprintf("%d", m.DepositNonce.Big().Int64()))
//			//w.log.Debug("voteProposal", "ResourceId", fmt.Sprintf("%x", m.ResourceId))
//			//tx, err := w.voteContract.VoteProposalBySigMachine(
//			//	w.conn.Opts(),
//			//	m.Source.Big(),
//			//	m.DepositNonce.Big(),
//			//	m.ResourceId,
//			//	dataHash,
//			//)
//			txHash, err := w.VoteProposalBySigMachine(m, dataHash)
//			if err != nil {
//				w.log.Info("voteProposal", "error", err)
//				return
//			}
//			w.conn.UnlockOpts()
//
//			if err == nil {
//				w.log.Info("Submitted proposal vote", "tx", txHash, "src", m.Source, "depositNonce", m.DepositNonce)
//				for i := 0; i < 25; i++ {
//					if w.proposalIsComplete(m, dataHash) {
//						w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
//						break
//					}
//					time.Sleep(time.Second * 2)
//				}
//				return
//			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
//				w.log.Debug("Nonce too low, will retry")
//				time.Sleep(TxRetryInterval)
//			} else {
//				w.log.Warn("Voting failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "err", err)
//				time.Sleep(TxRetryInterval)
//			}
//
//			if w.proposalIsComplete(m, dataHash) {
//				//w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
//				return
//			}
//		}
//	}
//	w.log.Error("Submission of Vote transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
//	w.sysErr <- ErrFatalTx
//}

//func (w *Writer) ExecuteProposal(m msg.Message, data []byte, dataHash [32]byte) {
//	w.muExec.Lock()
//	defer w.muExec.Unlock()
//
//	var status bool
//	var txHash string
//	var txHashRes string
//	receiveAt := time.Now().Format("2006-01-02 15:04:05")
//
//	defer func() {
//		if status {
//			model.UpdateExecuteStatus(m, 1, txHashRes, receiveAt)
//		}
//	}()
//
//	for i := 0; i < TxRetryLimit; i++ {
//		select {
//		case <-w.stop:
//			return
//		default:
//			_, eStatus, err := model.GetBridgeTxStatus(m)
//			if err == nil && eStatus > 0 {
//				w.log.Info("ExecuteProposal", "skip,src", m.Source, "depositNonce", m.DepositNonce)
//				return
//			}
//			//err = w.conn.LockAndUpdateOpts()
//			//if err != nil {
//			//	w.log.Error("Failed to update nonce", "err", err)
//			//	return
//			//}
//			//
//			//gasLimit := w.conn.Opts().GasLimit
//			//gasPrice := w.conn.Opts().GasPrice
//			//
//			//tx, err := w.voteContract.ExecuteProposal(
//			//	w.conn.Opts(),
//			//	m.Source.Big(),
//			//	m.DepositNonce.Big(),
//			//	data,
//			//)
//			//fmt.Println("~~~~~~~~ data ")
//			//fmt.Println(fmt.Sprintf("%x", data))
//			//w.conn.UnlockOpts()
//			txHashExec, txTime, err := w.ExecuteProposalBySigMachine(m, data)
//			if err != nil {
//				return
//			}
//
//			if err == nil {
//				txHash = txHashExec
//				w.log.Info("Submitted proposal execution", "tx", txHash, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
//				///
//				for j := 0; j < 5; j++ {
//					if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
//						status = true
//						txHashRes = txHash
//						receiveAt = txTime.Format("2006-01-02 15:04:05")
//						w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
//						break
//					}
//					time.Sleep(time.Second * 5)
//				}
//				///
//				return
//			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
//				w.log.Error("Nonce too low, will retry")
//				time.Sleep(TxRetryInterval)
//			} else {
//				w.log.Warn("Execution failed, proposal may already be complete", "err", err)
//				time.Sleep(TxRetryInterval)
//			}
//
//			if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
//				status = true
//				txHashRes = txHash
//				w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
//				return
//			}
//		}
//	}
//	w.log.Error("Submission of Execute transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
//	w.sysErr <- ErrFatalTx
//}

func (w *Writer) voteProposal(m msg.Message, dataHash [32]byte) (string, error) {

	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		return "", fmt.Errorf("failed to update tx opts err %v", err)
	}

	nonce := w.conn.Opts().Nonce
	gasLimit := w.conn.Opts().GasLimit
	gasPrice := w.conn.Opts().GasPrice
	inputData, err := signature.GenerateVoteProposalInputData(big.NewInt(int64(m.Source)), big.NewInt(int64(m.DepositNonce)), m.ResourceId, dataHash)
	if err != nil {
		return "", err
	}
	to := common.HexToAddress(w.Cfg.VoteContractAddress)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce.Uint64(),
		To:       &to,
		Value:    nil,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     inputData,
	})

	apiSecret := os.Getenv("API_SECRET_SIG_MACHINE")
	if len(apiSecret) <= 0 {
		apiSecret = "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	}
	chainId, err := w.conn.connEvm.ChainID(context.Background())
	if err != nil {
		return "", fmt.Errorf("get chain id err %v", err)
	}

	err = signature.SignAndSendTxEth(w.conn.connEvm, chainId.Uint64(), tx, apiSecret)
	if err != nil {
		return "", fmt.Errorf("vote propose err %v", err)
	}

	return tx.Hash().String(), nil

}

func (w *Writer) ExecuteProposal(m msg.Message, data []byte, dataHash [32]byte) (string, *time.Time, error) {

	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		return "", nil, fmt.Errorf("failed to update tx opts err %v", err)
	}

	nonce := w.conn.Opts().Nonce
	gasLimit := w.conn.Opts().GasLimit
	gasPrice := w.conn.Opts().GasPrice
	inputData, err := signature.GenerateExecuteProposalInputData(big.NewInt(int64(m.Source)), big.NewInt(int64(m.DepositNonce)), data)
	if err != nil {
		return "", nil, err
	}
	to := common.HexToAddress(w.Cfg.VoteContractAddress)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce.Uint64(),
		To:       &to,
		Value:    nil,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     inputData,
	})
	tx.Time()

	apiSecret := os.Getenv("API_SECRET_SIG_MACHINE")
	if len(apiSecret) <= 0 {
		apiSecret = "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	}
	chainId, err := w.conn.connEvm.ChainID(context.Background())
	if err != nil {
		return "", nil, fmt.Errorf("get chain id err %v", err)
	}
	err = signature.SignAndSendTxEth(w.conn.connEvm, chainId.Uint64(), tx, apiSecret)
	if err != nil {
		return "", nil, fmt.Errorf("vote propose err %v", err)
	}
	txTime := tx.Time()
	return tx.Hash().String(), &txTime, nil

}
