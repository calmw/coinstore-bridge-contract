package event

import (
	"coinstore/bridge/config"
	"coinstore/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/status-im/keycard-go/hexutils"
	"io"
	"net/http"
	"strings"
)

type EventLog struct {
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Num0               string `json:"0"`
			Num1               string `json:"1"`
			Num2               string `json:"2"`
			Num3               string `json:"3"`
			Num4               string `json:"4"`
			Num5               string `json:"5"`
			Amount0In          string `json:"amount0In"`
			Amount1In          string `json:"amount1In"`
			Amount1Out         string `json:"amount1Out"`
			Sender             string `json:"sender"`
			Amount0Out         string `json:"amount0Out"`
			To                 string `json:"to"`
			ResourceID         string `json:"resourceID"`
			DepositNonce       string `json:"depositNonce"`
			DestinationChainId string `json:"destinationChainId"`
			Data               string `json:"data"`
		} `json:"result"`
		ResultType struct {
			Amount0In  string `json:"amount0In"`
			Amount1In  string `json:"amount1In"`
			Amount1Out string `json:"amount1Out"`
			Sender     string `json:"sender"`
			Amount0Out string `json:"amount0Out"`
			To         string `json:"to"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionID string `json:"transaction_id"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type EventData struct {
	Data               string `json:"data"`
	TxHash             string `json:"tx_hash"`
	ResourceID         string `json:"resourceID"`
	DepositNonce       string `json:"depositNonce"`
	DestinationChainId string `json:"destinationChainId"`
}

type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  []Param     `json:"params,omitempty"`
	ID      interface{} `json:"id"`
}

type Param struct {
	Address   []string `json:"address"`
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Topics    []string `json:"topics"`
}

func GetEventData(number int64) ([]EventData, error) {
	var result []EventData
	url := fmt.Sprintf("%s/v1/blocks/%d/events", config.TronApiHost, number)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var eventLog EventLog
	err := json.Unmarshal(body, &eventLog)
	if err != nil || !eventLog.Success {
		return nil, errors.New("failed to parse event log")
	}
	for _, d := range eventLog.Data {
		if d.EventName == "Deposit" {
			result = append(result, EventData{
				Data:               d.Result.Data,
				TxHash:             d.TransactionID,
				ResourceID:         d.Result.ResourceID,
				DepositNonce:       d.Result.DepositNonce,
				DestinationChainId: d.Result.DestinationChainId,
			})
		}
	}

	return result, nil
}

type EthCallJsonRpcRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  string `json:"params"`
	ID      int    `json:"id"`
}

type JsonRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

func GetDepositRecord(from, to, data string) (utils.DepositRecord, error) {
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return utils.DepositRecord{}, err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return utils.DepositRecord{}, err
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
}`, from, to, data, utils.RandInt(100, 10000))
	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	err := json.Unmarshal(body, &jsonRpcResponse)
	if err != nil {
		return utils.DepositRecord{}, errors.New("eth call failed")
	}
	return utils.ParseBridgeDepositRecordData(hexutils.HexToBytes("197649b0" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}

func ResourceIdToTokenInfo(from, to, data string) (utils.TokenInfo, error) {
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	if !strings.HasPrefix(from, "0x") {
		fromAddress, err := address.Base58ToAddress(from)
		if err != nil {
			return utils.TokenInfo{}, err
		}
		from = fromAddress.Hex()
	}
	if !strings.HasPrefix(to, "0x") {
		toAddress, err := address.Base58ToAddress(to)
		if err != nil {
			return utils.TokenInfo{}, err
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
}`, from, to, data, utils.RandInt(100, 10000))

	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var jsonRpcResponse JsonRpcResponse
	err := json.Unmarshal(body, &jsonRpcResponse)
	if err != nil {
		return utils.TokenInfo{}, errors.New("eth call failed")
	}
	return utils.ParseBridgeResourceIdToTokenInfo(hexutils.HexToBytes("6cbfe81f" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))
}
