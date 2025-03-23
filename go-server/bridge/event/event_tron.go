package event

import (
	"coinstore/bridge/config"
	"coinstore/bridge/msg"
	"coinstore/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type EvtLog struct {
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

type EvtData struct {
	Data               string `json:"data"`
	TxHash             string `json:"tx_hash"`
	ResourceID         string `json:"resourceID"`
	DepositNonce       string `json:"depositNonce"`
	DestinationChainId string `json:"destinationChainId"`
}

type ProposalVoteEvent struct {
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Num0          string `json:"0"`
			Num1          string `json:"1"`
			Num2          string `json:"2"`
			Num3          string `json:"3"`
			Num4          string `json:"4"`
			DepositNonce  string `json:"depositNonce"`
			ResourceID    string `json:"resourceID"`
			DataHash      string `json:"dataHash"`
			OriginChainID string `json:"originChainID"`
			Status        string `json:"status"`
		} `json:"result"`
		ResultType struct {
			DepositNonce  string `json:"depositNonce"`
			ResourceID    string `json:"resourceID"`
			DataHash      string `json:"dataHash"`
			OriginChainID string `json:"originChainID"`
			Status        string `json:"status"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionID string `json:"transaction_id"`
		Unconfirmed   bool   `json:"_unconfirmed"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type ProposalEventData struct {
	TxHash             string   `json:"tx_hash"`
	ResourceID         [32]byte `json:"resourceID"`
	OriginDepositNonce uint64   `json:"originDepositNonce"`
	OriginChainID      int64    `json:"originChainID"`
	DataHash           [32]byte `json:"dataHash"`
	ProposalStatus     int64    `json:"proposalStatus"`
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

type EthCallJsonRpcRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  string `json:"params"`
	ID      int    `json:"id"`
}

func GetEventData(number int64) ([]EvtData, error) {
	var result []EvtData
	url := fmt.Sprintf("%s/v1/blocks/%d/events", config.TronApiHost, number)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var eventLog EvtLog
	err := json.Unmarshal(body, &eventLog)
	if err != nil || !eventLog.Success {
		return nil, errors.New("failed to parse event log")
	}
	for _, d := range eventLog.Data {
		if d.EventName == "Deposit" {
			result = append(result, EvtData{
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

func ParseProposalEvent(originChainId msg.ChainId, originDepositNonce msg.Nonce, number int64) ([]ProposalEventData, error) {
	var result []ProposalEventData
	url := fmt.Sprintf("%s/v1/blocks/%d/events", config.TronApiHost, number)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("---=====", number)
	fmt.Println("---=====", string(body))
	var eventLog ProposalVoteEvent
	fmt.Println("---=====-------------------------------------------------------------------------------------------------")
	err := json.Unmarshal(body, &eventLog)
	if err != nil || !eventLog.Success {
		return nil, errors.New("failed to parse event log")
	}
	for _, d := range eventLog.Data {
		nonce, err := strconv.ParseInt(d.Result.DepositNonce, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse event log %v", err)
		}
		chainID, err := strconv.ParseInt(d.Result.OriginChainID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse event log %v", err)
		}
		fmt.Println(d.EventName, nonce, int64(originDepositNonce), chainID, int64(originChainId))
		if (d.EventName == "ProposalEvent") && (nonce == int64(originDepositNonce)) && (chainID == int64(originChainId)) {
			fmt.Println(d)
			fmt.Println(d.Result.DepositNonce)
			result = append(result, ProposalEventData{
				//TxHash:             d.TransactionID,
				//ResourceID:         d.Result.ResourceID,
				//OriginDepositNonce: d.Result.DepositNonce,
				//OriginChainID:      d.Result.OriginChainID,
				//DataHash:           d.Result.DataHash,
				//ProposalStatus:     d.Result.Status,
			})
		}
	}

	return result, nil
}

func GetProposalEvent(destChainId msg.ChainId, depositNonce msg.Nonce, number int64) ([]ProposalEventData, error) {
	event, err := ParseProposalEvent(destChainId, depositNonce, number)
	if err != nil {
		return nil, err
	}
	return event, nil
}

type EthTxData struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash        string      `json:"blockHash"`
		BlockNumber      string      `json:"blockNumber"`
		From             string      `json:"from"`
		Gas              string      `json:"gas"`
		GasPrice         string      `json:"gasPrice"`
		Hash             string      `json:"hash"`
		Input            string      `json:"input"`
		Nonce            interface{} `json:"nonce"`
		R                string      `json:"r"`
		S                string      `json:"s"`
		To               string      `json:"to"`
		TransactionIndex string      `json:"transactionIndex"`
		Type             string      `json:"type"`
		V                string      `json:"v"`
		Value            string      `json:"value"`
	} `json:"result"`
}

func EthGetTransactionByHash(txHash string) (int64, error) {
	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	ethCallBody := fmt.Sprintf(`{
	"jsonrpc": "2.0",
	"method": "eth_getTransactionByHash",
	"params": ["%s"],
	"id": %d
}`, txHash, utils.RandInt(100, 10000))
	fmt.Println(ethCallBody)
	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var ethTxData EthTxData
	err := json.Unmarshal(body, &ethTxData)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(body))
	fmt.Println("==")
	num, err := strconv.ParseInt(ethTxData.Result.BlockNumber, 10, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		return 0, err
	}

	return num, nil
}
