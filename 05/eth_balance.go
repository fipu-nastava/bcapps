package main

import (
	"encoding/json"
	"net/http"
	"io"
	"bytes"
	"fmt"
	"math/big"
)

type JsonRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []string `json:"params"`
	Id int `json:"id"`
}
type JsonBalanceResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result string `json:"result"`
	Id int `json:"id"`
}
type JsonAccountsResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result []string `json:"result"`
	Id int `json:"id"`
}

func NewRequest(method string, params []string) (r *JsonRequest) {
	r = new(JsonRequest)
	r.Jsonrpc = "2.0"
	r.Method = method
	r.Params = params
	r.Id = 1
	return
}
func (r *JsonRequest) ToReader() (rdr io.Reader) {
	b, err := json.Marshal(r)
	if err != nil {
		panic("Greška u JSON")
		return
	}

	rdr = bytes.NewReader(b)
	return
}

func ListAccounts(url string) []string {
	a := NewRequest("eth_accounts", []string {})
	r, err := http.Post(url, "application/json", a.ToReader())
	if err != nil {
		panic(err)
	}
	respDecoder := json.NewDecoder(r.Body)
	result := new(JsonAccountsResult)
	respDecoder.Decode(result)
	return result.Result
}
func GetBalance(url string, account string) float64 {
	req := NewRequest("eth_getBalance", []string {account, "latest"})
	r, err := http.Post(url, "application/json", req.ToReader())
	if err != nil {
		panic(err)
	}
	respDecoder := json.NewDecoder(r.Body)
	result := new(JsonBalanceResult)
	respDecoder.Decode(result)

	bal := result.Result
	a, _ := new(big.Int).SetString(bal, 0)
	e := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	balance := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(e))

	retval, _ := balance.Float64()
	return retval
}
func GetTxCount(url string, account string) uint64 {
	req := NewRequest("eth_getTransactionCount", []string {account, "latest"})
	r, err := http.Post(url, "application/json", req.ToReader())
	if err != nil {
		panic(err)
	}
	respDecoder := json.NewDecoder(r.Body)
	result := new(JsonBalanceResult)
	respDecoder.Decode(result)

	bal := result.Result
	a, _ := new(big.Int).SetString(bal, 0)
	return a.Uint64()
}


func main() {
	url := "http://localhost:7545"

	accounts := ListAccounts(url)
	for _, a := range accounts {
		balance := GetBalance(url, a)
		nonce := GetTxCount(url, a)
		fmt.Printf("Account %s has %f ETH (%d txs)\n", a, balance, nonce)
	}
}
