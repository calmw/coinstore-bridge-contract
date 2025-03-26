package jsonrpc

import (
	"coinstore/utils"
	"fmt"
	log "github.com/calmw/clog"
	"testing"
)

func Test_jsonRpc(t *testing.T) {
	r := NewJsonRpc("https://api.shasta.trongrid.io/jsonrpc", log.Root())
	res, err := r.Call("eth_blockNumber", []interface{}{})
	fmt.Println(err, string(res))
	fmt.Println(utils.HexToBigInt(string(res)))
}
