package jsonrpc

import (
	"fmt"
	log "github.com/calmw/blog"
	"testing"
)

func Test_jsonRpc(t *testing.T) {
	r := NewJsonRpc("https://api.shasta.trongrid.io/jsonrpc", log.Root())
	res, err := r.Call("eth_chainId", []interface{}{})
	fmt.Println(err, string(res))
}
