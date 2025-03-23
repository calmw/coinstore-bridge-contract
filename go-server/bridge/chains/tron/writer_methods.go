package tron

import (
	"coinstore/bridge/msg"
	"coinstore/model"
	"coinstore/utils"
	"errors"
	log "github.com/calmw/clog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"math/big"
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
	prop, err := w.voteContract.GetProposal(m.Source.Big(), m.DepositNonce.Big(), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	///
	w.log.Debug("🫥  proposalIsComplete", "status", prop.Status, "nonce", m.DepositNonce)
	if prop.Status >= 2 {
		model.UpdateVoteStatus(m, 1)
	}
	///
	return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
}

func (w *Writer) proposalIsFinalized(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.voteContract.GetProposal(srcId.Big(), nonce.Big(), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
}

func (w *Writer) proposalIsPassed(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.voteContract.GetProposal(srcId.Big(), nonce.Big(), dataHash)
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
	from, err := address.Base58ToAddress(w.Cfg.From)
	if err != nil {
		panic(err)
	}
	hasVoted, err := w.voteContract.HasVotedOnProposal(IDAndNonce(srcId, nonce), dataHash, common.HexToAddress(from.Hex()))
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
	w.log.Info("Creating generic proposal", "src", m.Source, "nonce", m.DepositNonce)

	metadata := m.Payload[0].([]byte)
	data := ConstructGenericProposalData(metadata)
	bridgeAddress, err := address.Base58ToAddress(w.Cfg.BridgeContractAddress)
	toHash := append(common.HexToAddress(bridgeAddress.Hex()).Bytes(), data...)
	dataHash := utils.Hash(toHash)
	//fmt.Printf("####---%x\n", dataHash)

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
			w.log.Trace("No finalization event found in current block", "block", latestBlock, "src", m.Source, "nonce", m.DepositNonce)
			latestBlock = latestBlock.Add(latestBlock, big.NewInt(1))
		}
	}
	log.Warn("Block watch limit exceeded, skipping execution", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
}

func (w *Writer) voteProposal(m msg.Message, dataHash [32]byte) {
	w.muVote.Lock()
	defer w.muVote.Unlock()

	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			//err := w.conn.LockAndUpdateOpts()
			//if err != nil {
			//	w.log.Error("Failed to update tx opts", "err", err)
			//	continue
			//}

			//gasLimit := w.conn.Opts().GasLimit
			//gasPrice := w.conn.Opts().GasPrice

			txHash, err := w.voteContract.VoteProposal(
				m.Source.Big(),
				m.DepositNonce.Big(),
				m.ResourceId,
				dataHash,
			)
			//w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Submitted proposal vote", "tx", txHash, "src", m.Source, "depositNonce", m.DepositNonce)
				for i := 0; i < 25; i++ {
					if w.proposalIsComplete(m, dataHash) {
						w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
						break
					}
					time.Sleep(time.Second * 2)
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Voting failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "err", err)
				time.Sleep(TxRetryInterval)
			}

			if w.proposalIsComplete(m, dataHash) {
				//w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Vote transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

func (w *Writer) ExecuteProposal(m msg.Message, data []byte, dataHash [32]byte) {
	w.muExec.Lock()
	defer w.muExec.Unlock()

	var err error
	var status bool
	var txHash string
	var txHashRes string
	receiveAt := time.Now().Format("2006-01-02 15:04:05")

	defer func() {
		if status {
			model.UpdateExecuteStatus(m, 1, txHashRes, receiveAt)
		}
	}()

	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			txHash, err = w.voteContract.ExecuteProposal(
				m.Source.Big(),
				m.DepositNonce.Big(),
				data,
				m.ResourceId,
			)

			if err == nil {
				w.log.Info("Submitted proposal execution", "tx", txHash, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				///
				for j := 0; j < 5; j++ {
					if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
						status = true
						txHashRes = txHash
						receiveAt = time.Now().Format("2006-01-02 15:04:05")
						w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
						break
					}
					time.Sleep(time.Second * 5)
				}
				///
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Error("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Execution failed, proposal may already be complete", "err", err)
				time.Sleep(TxRetryInterval)
			}

			if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
				status = true
				w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Execute transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}
