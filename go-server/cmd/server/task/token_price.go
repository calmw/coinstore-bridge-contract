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
	fmt.Println("~")
	resp := &http.Response{}
	var res BinancePrice
	defer resp.Body.Close()
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
		if token == "ETH" {
			TokenPrice["binance"]["ETH"] = res.Price
			TokenPrice["binance"]["TETH"] = res.Price
			TokenPrice["binance"]["WETH"] = res.Price
		}
	}

}

func GetPrice(symbol string) string {
	return TokenPrice["binance"][symbol]
}
