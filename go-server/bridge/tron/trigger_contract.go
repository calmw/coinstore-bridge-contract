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
	return ParseVoteGetProposal(hexutils.HexToBytes("5b95771f" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func CallBody(from, to string, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (string, error) {
	requestData, err := GenerateVoteHasVotedOnProposal(arg0, arg1, arg2)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return "", err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return "", err
		}
		to = toAddress.Hex()
	}
	ethCallBody := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_call","params":[{"from":"%s","to":"%s","gas":"0x0","gasPrice":"0x0","value":"0x0","data":"%s"},"latest"],"id": %d}`,
		from, to, requestData, utils.RandInt(100, 10000))
	fmt.Println(ethCallBody)
	return ethCallBody, nil
}
func HasVotedOnProposal(from, to string, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	ethCallBody, err := CallBody(from, to, arg0, arg1, arg2)
	fmt.Println(ethCallBody)
	fmt.Println(err)
	//	ethCallBody = `{
	//	"jsonrpc": "2.0",
	//	"method": "eth_call",
	//	"params": [{
	//		"from": "0x41f4a0d088ef4ec7b0e231cc365f16726ad552e051",
	//		"to": "0x41f4a0d088ef4ec7b0e231cc365f16726ad552e051",
	//		"gas": "0x0",
	//		"gasPrice": "0x0",
	//		"value": "0x0",
	//		"data": "0xc70bf0b50000000000000000000000000000000000000000000000000000000000000f02f5ffe0a02a4b931713566e54bfafc450192c48bc69010f471fa6dd2d2639a65b0000000000000000000000003942fda93c573e2ce9e85b0bb00ba98a144f27f6"
	//	}, "latest"],
	//	"id": 8400
	//}`
	//	fmt.Println(ethCallBody)
	req, _ := http.NewRequest("POST", "https://nile.trongrid.io/jsonrpc", strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	//err = json.Unmarshal(body, &jsonRpcResponse)
	//if err != nil {
	//	return false, errors.New("eth call failed")
	//}
	//fmt.Println(jsonRpcResponse.Result)
	fmt.Println("==")
	fmt.Println(string(body))
	fmt.Println("==")
	fmt.Println(jsonRpcResponse)
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
	return hexutils.BytesToHex(tx.GetTxid()), nil
}

func ExecuteProposal(cli *client.GrpcClient, from, contractAddress string, ks *keystore.KeyStore, ka *keystore.Account, originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (string, error) {
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes\":\"%s\"},{\"bytes32\":\"%s\"}]",
		originChainId.String(), originDepositNonce.String(), hexutils.BytesToHex(data[:]), hexutils.BytesToHex(resourceId[:]),
	)
	fmt.Println("--")
	fmt.Println(triggerData)

	tx, err := cli.TriggerContract(from, contractAddress, "voteProposal(uint256,uint256,bytes32,bytes32)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	return hexutils.BytesToHex(tx.GetTxid()), nil
}
