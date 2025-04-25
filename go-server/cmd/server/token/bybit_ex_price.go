package token

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
	"log"
)

func GetBybitPrice() {
	client := bybit.NewBybitHttpClient("tantinchain", "bGGmXFN8575wnlVYF2", bybit.WithBaseURL(bybit.MAINNET))
	for _, token := range tokens {
		symbol := fmt.Sprintf("%sUSDT", token)
		params := map[string]interface{}{"category": "spot", "symbol": symbol}
		tickers, err := client.NewUtaBybitServiceWithParams(params).GetMarketTickers(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		m, ok := tickers.Result.(map[string]interface{})
		if !ok {
			log.Println("类型错误")
			return
		}
		list, ok := m["list"].([]interface{})
		if !ok {
			log.Println("类型错误")
			return
		}
		if len(list) <= 0 {
			log.Println("数据错误")
			return
		}
		priceMap, ok := list[0].(map[string]interface{})
		if !ok {
			log.Println("类型错误")
			return
		}
		priceIntf, ok := priceMap["lastPrice"]
		if !ok {
			log.Println("类型错误")
			return
		}
		price, ok := priceIntf.(string)
		if !ok {
			log.Println("类型错误")
			return
		}
		ExTokenPriceData.Set(ExBibit, token, price)
	}

}
