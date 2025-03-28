package tron

import (
	"coinstore/bridge/config"
	"coinstore/utils"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	tcommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/status-im/keycard-go/hexutils"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
)

func GetDepositRecord(from, to string, destinationChainId, depositNonce *big.Int) (DepositRecord, error) {
	requestData, err := GenerateBridgeDepositRecordsData(destinationChainId, depositNonce)
	if err != nil {
		return DepositRecord{}, fmt.Errorf("generateBridgeDepositRecordsData error %v", err)
	}
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return DepositRecord{}, err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return DepositRecord{}, err
		}
		to = toAddress.Hex()
	}
	ethCallBody := fmt.Sprintf(`{
	"jsonrpc": "2.0",
	"method": "eth_call",
	"params": [{
		"from": "%s",
		"to": "%s",
		"gas": "0x0",
		"gasPrice": "0x0",
		"value": "0x0",
		"data": "%s"
	}, "latest"],
	"id": %d
}`, from, to, requestData, utils.RandInt(100, 10000))
	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	err = json.Unmarshal(body, &jsonRpcResponse)
	if err != nil {
		return DepositRecord{}, errors.New("eth call failed")
	}
	return ParseBridgeDepositRecordData(hexutils.HexToBytes("197649b0" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func ResourceIdToTokenInfo(from, to string, requestId [32]byte) (TokenInfo, error) {
	requestData, err := GenerateBridgeGetTokenInfoByResourceId(requestId)
	if err != nil {
		return TokenInfo{}, err
	}
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return TokenInfo{}, err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return TokenInfo{}, err
		}
		to = toAddress.Hex()
	}
	ethCallBody := fmt.Sprintf(`{
	"jsonrpc": "2.0",
	"method": "eth_call",
	"params": [{
		"from": "%s",
		"to": "%s",
		"gas": "0x0",
		"gasPrice": "0x0",
		"value": "0x0",
		"data": "%s"
	}, "latest"],
	"id": %d
}`, from, to, requestData, utils.RandInt(100, 10000))
	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	err = json.Unmarshal(body, &jsonRpcResponse)
	if err != nil {
		return TokenInfo{}, errors.New("eth call failed")
	}
	return ParseBridgeResourceIdToTokenInfo(hexutils.HexToBytes("6cbfe81f" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func GetProposal(from, to string, originChainID *big.Int, depositNonce *big.Int, dataHash [32]byte) (IVoteProposal, error) {
	requestData, err := GenerateVoteGetProposal(originChainID, depositNonce, dataHash)
	if err != nil {
		return IVoteProposal{}, err
	}
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return IVoteProposal{}, err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return IVoteProposal{}, err
		}
		to = toAddress.Hex()
	}
	ethCallBody := fmt.Sprintf(`{
	"jsonrpc": "2.0",
	"method": "eth_call",
	"params": [{
		"from": "%s",
		"to": "%s",
		"gas": "0x0",
		"gasPrice": "0x0",
		"value": "0x0",
		"data": "%s"
	}, "latest"],
	"id": %d
}`, from, to, requestData, utils.RandInt(100, 10000))
	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	err = json.Unmarshal(body, &jsonRpcResponse)
	if err != nil {
		return IVoteProposal{}, errors.New("eth call failed")
	}
	return ParseVoteGetProposal(hexutils.HexToBytes("5b95771f" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func HasVotedOnProposal(from, to string, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	requestData, err := GenerateVoteHasVotedOnProposal(arg0, arg1, arg2)
	if err != nil {
		return false, err
	}
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return false, err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return false, err
		}
		to = toAddress.Hex()
	}
	ethCallBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_call","params":[{"from":"%s","to":"%s","gas":"0x0","gasPrice":"0x0","value":"0x0","data":"%s"},"latest"],"id": %d}`,
		from, to, requestData, utils.RandInt(100, 10000))
	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	err = json.Unmarshal(body, &jsonRpcResponse)
	if err != nil {
		return false, errors.New("eth call failed")
	}
	return ParseVoteHasVotedOnProposal(hexutils.HexToBytes("c70bf0b5" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func VoteProposal(cli *client.GrpcClient, from, contractAddress string, ks *keystore.KeyStore, ka *keystore.Account, originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (string, error) {
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes32\":\"%s\"},{\"bytes32\":\"%s\"}]",
		originChainId.String(), originDepositNonce.String(), hexutils.BytesToHex(resourceId[:]), hexutils.BytesToHex(dataHash[:]),
	)
	tx, err := cli.TriggerContract(from, contractAddress, "voteProposal(uint256,uint256,bytes32,bytes32)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	tx.GetLogs()
	return hexutils.BytesToHex(tx.GetTxid()), nil
}

const (
	AccountName = "my_account"
	Passphrase  = "account_pwd"
)

func ExecuteProposal(cli *client.GrpcClient, from, contractAddress string, ks *keystore.KeyStore, ka *keystore.Account, originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (string, error) {
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes\":\"%s\"},{\"bytes32\":\"%s\"}]",
		originChainId.String(), originDepositNonce.String(), hexutils.BytesToHex(data[:]), hexutils.BytesToHex(resourceId[:]),
	)

	privateKey := os.Getenv("COIN_STORE_BRIDGE_TRON")
	_, _, err := getKeyFromPrivateKey(privateKey, AccountName, Passphrase)
	//if strings.Contains(err.Error(),"already exists")
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return "", err
	}
	// 获得keystore与account
	ks, ka, err = unlockedKeystore(from, Passphrase)
	if err != nil {
		return "", err
	}

	tx, err := cli.TriggerContract(from, contractAddress, "executeProposal(uint256,uint256,bytes,bytes32)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	return hexutils.BytesToHex(tx.GetTxid()), nil
}

func getKeyFromPrivateKey(privateKey, name, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	privateKey = strings.TrimPrefix(privateKey, "0x")

	if store.DoesNamedAccountExist(name) {
		return nil, nil, fmt.Errorf("account %s already exists", name)
	}

	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, nil, err
	}
	if len(privateKeyBytes) != tcommon.Secp256k1PrivateKeyBytesLength {
		return nil, nil, tcommon.ErrBadKeyLength
	}

	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	ks := store.FromAccountName(name)
	account, err := ks.ImportECDSA(sk.ToECDSA(), passphrase)
	if err != nil {
		return nil, nil, err
	}
	return store.FromAccountName(name), &account, err
}

func unlockedKeystore(from, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	sender, err := address.Base58ToAddress(from)
	if err != nil {
		return nil, nil, fmt.Errorf("address not valid: %s", from)
	}
	ks := store.FromAddress(from)
	if ks == nil {
		return nil, nil, fmt.Errorf("could not open local keystore for %s", from)
	}
	account, lookupErr := ks.Find(keystore.Account{Address: sender})
	if lookupErr != nil {
		return nil, nil, fmt.Errorf("could not find %s in keystore", from)
	}
	if unlockError := ks.Unlock(account, passphrase); unlockError != nil {
		return nil, nil, errors.Unwrap(unlockError)
	}
	return ks, &account, nil
}
