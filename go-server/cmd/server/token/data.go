package token

import (
	"github.com/shopspring/decimal"
	"log"
	"strings"
	"sync"
)

const (
	ExBinance = "binance" // 币安
	ExBibit   = "bybit"   // Bybit
)

var ExTokenPriceData *PriceData
var tokens = []string{"ETH", "USDC"}

type PriceData struct {
	rw    sync.RWMutex
	price map[string]string            // 最终价格
	data  map[string]map[string]string // 各个交易所价格数据
}

func NewTokenPrice() *PriceData {
	return &PriceData{
		rw:    sync.RWMutex{},
		price: make(map[string]string),
		data:  make(map[string]map[string]string),
	}
}

func InitPriceData() {
	ExTokenPriceData = NewTokenPrice()
}

func (m *PriceData) Set(exchangeName, tokenName, tokenPrice string) {
	m.rw.Lock()
	defer m.rw.Unlock()
	if strings.Contains(tokenName, "ETH") {
		tokenName = "ETH"
	}
	log.Println("set price ", exchangeName, tokenName, tokenPrice)
	if m.data[exchangeName] == nil {
		m.data[exchangeName] = map[string]string{}
	}
	m.data[exchangeName][tokenName] = tokenPrice
	exchange2Name := ExBibit
	if exchangeName == ExBibit {
		exchange2Name = ExBinance
	}
	if m.data[exchange2Name] == nil {
		m.data[exchange2Name] = map[string]string{}
		log.Print("1")
		return
	}
	exchange2Price := m.data[exchange2Name][tokenName]
	if len(exchange2Price) == 0 {
		m.price[tokenName] = tokenPrice
		return
	}
	price2Deci, err := decimal.NewFromString(exchange2Price)
	if err != nil {
		return
	}
	priceDeci, err := decimal.NewFromString(tokenPrice)
	if err != nil {
		return
	}
	priceDeci = priceDeci.Add(price2Deci).Div(decimal.NewFromInt(2))
	m.price[tokenName] = priceDeci.String()
}

func (m *PriceData) Get(tokenName string) (string, bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	tokenName = strings.ToUpper(tokenName)
	if tokenName == "USDT" {
		return "1", true
	}
	if strings.Contains(tokenName, "ETH") {
		tokenName = "ETH"
	}
	if strings.Contains(tokenName, "USDC") {
		tokenName = "USDC"
	}
	s, ok := m.price[tokenName]
	return s, ok
}
