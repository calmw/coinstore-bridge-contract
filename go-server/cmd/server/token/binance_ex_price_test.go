package token

import "testing"

func TestGetBinancePrice(t *testing.T) {
	InitPriceData()
	GetBinancePrice()
}
