package tron

import (
	"coinstore/bridge/config"
	"coinstore/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/status-im/keycard-go/hexutils"
	"io"
	"log"
	"math/big"
	"net/http"
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
	fmt.Println("~")
	fmt.Println(jsonRpcResponse.Result)
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
		return false, errors.New("eth call failed")
	}
	return ParseVoteHasVotedOnProposal(hexutils.HexToBytes("6cbfe81f" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func VoteProposal(cli *client.GrpcClient, from, contractAddress string, ks *keystore.KeyStore, ka *keystore.Account, originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (string, error) {
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes32\":\"%s\"},{\"bytes32\":\"%s\"}]",
		originChainId.String(), originDepositNonce.String(), hexutils.BytesToHex(resourceId[:]), hexutils.BytesToHex(dataHash[:]),
	)
	fmt.Println(triggerData)

	tx, err := cli.TriggerContract(from, contractAddress, "voteProposal(uint256,uint256,bytes32,bytes32)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", tx.Txid)
	return hexutils.BytesToHex(tx.GetTxid()), nil
}

func ExecuteProposal(cli *client.GrpcClient, from, contractAddress string, ks *keystore.KeyStore, ka *keystore.Account, originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (string, error) {
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes\":\"%s\"},{\"bytes32\":\"%s\"}]",
		originChainId.String(), originDepositNonce.String(), hexutils.BytesToHex(data[:]), hexutils.BytesToHex(resourceId[:]),
	)
	fmt.Println(triggerData)

	tx, err := cli.TriggerContract(from, contractAddress, "voteProposal(uint256,uint256,bytes32,bytes32)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", tx.Txid)
	return hexutils.BytesToHex(tx.GetTxid()), nil
}
