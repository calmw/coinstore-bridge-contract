package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/calmw/clog"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type RpcRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      interface{} `json:"id,omitempty"`
}

type RpcResponse struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   interface{}     `json:"error"`
	ID      interface{}     `json:"id"`
}

type JsonRpc struct {
	Uri string `json:"uri"`
	log log.Logger
}

func NewJsonRpc(uri string, log log.Logger) *JsonRpc {
	return &JsonRpc{Uri: uri, log: log}
}

func (j *JsonRpc) Call(method string, params interface{}) (json.RawMessage, error) {
	request := RpcRequest{
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000) + 1,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshalling error:%v", err)
	}
	log.Debug("jsonrpc", "post", strings.Replace(string(data), `"`, "'", -1))

	resp, err := http.Post(j.Uri, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("sending request error:%v", err)
	}
	defer resp.Body.Close()

	var response RpcResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decoding response error:%v", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("jsonrpc error:%v", response.Error)
	}

	return response.Result, nil
}
