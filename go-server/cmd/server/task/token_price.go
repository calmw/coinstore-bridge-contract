package task

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var TokenPrice = map[string]map[string]string{}

type BinancePrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

var tokens = []string{"ETH", "USDC"}

func GetBinancePrice() {
	resp := &http.Response{}
	var res BinancePrice
	var err error
	for _, token := range tokens {
		url := fmt.Sprintf("https://api2.binance.com/api/v3/ticker/price?symbol=%sUSDT", token)
		// 发送HTTP GET请求
		resp, err = http.Get(url)
		if err != nil {
			log.Println(err)
			continue
		}

		// 读取响应体
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Println(err)
			continue
		}
		if TokenPrice["binance"] == nil {
			TokenPrice["binance"] = map[string]string{}
		}
		if token == "ETH" {
			TokenPrice["binance"]["ETH"] = res.Price
			TokenPrice["binance"]["TETH"] = res.Price
			TokenPrice["binance"]["WETH"] = res.Price
		} else {
			TokenPrice["binance"][token] = res.Price
		}
	}

	resp.Body.Close()

}

func GetPrice(symbol string) string {
	return TokenPrice["binance"][symbol]
}
