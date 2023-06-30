package restgo

import (
	"testing"
)

func Test_Get(t *testing.T) {
	type Data struct {
		HeadSlot string `json:"head_slot"`
	}
	type Response struct {
		Data Data `json:"data"`
	}
	var res Response
	err := Get("http://35.87.169.9:5052/eth/v1/node/syncing", &res)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("res: %+v\n", res)
}

/*

request:

curl https://docs-demo.quiknode.pro/ \
-X POST \
-H "Content-Type: application/json" \
--data '{"method":"eth_chainId","params":[],"id":1,"jsonrpc":"2.0"}'


response:

{"jsonrpc":"2.0","id":1,"result":"0x1"}

*/

func Test_Post(t *testing.T) {
	type Request struct {
		Method  string   `json:"method"`
		Params  []string `json:"params"`
		Id      int      `json:"id"`
		Jsonrpc string   `json:"jsonrpc"`
	}
	type Response struct {
		Id      int    `json:"id"`
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
	}
	var res Response
	err := Post("https://docs-demo.quiknode.pro", &Request{
		Method:  "eth_chainId",
		Params:  nil,
		Id:      1,
		Jsonrpc: "2.0",
	}, &res)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("res: %+v\n", res)
}
