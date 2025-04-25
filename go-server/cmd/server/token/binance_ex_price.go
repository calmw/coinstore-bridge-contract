package token

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BinancePrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

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
		ExTokenPriceData.Set(ExBinance, token, res.Price)
	}
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
