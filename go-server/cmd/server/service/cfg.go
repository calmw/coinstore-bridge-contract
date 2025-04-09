package service

import (
	"coinstore/db"
	"coinstore/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChainInfo struct {
	ChainId               uint64  `json:"chain_id"`
	BlockChainId          uint64  `json:"block_chain_id"`
	BridgeContractAddress string  `json:"bridge_contract_address"`
	Endpoint              string  `json:"endpoint"`
	ChainName             string  `json:"chain_name"`
	Explorer              string  `json:"explorer"`
	Logo                  string  `json:"logo"`
	NativeCoinName        string  `json:"native_coin_name"`
	NativeCoinSymbol      string  `json:"native_coin_symbol"`
	NativeCoinDecimals    string  `json:"native_coin_decimals"`
	Tokens                []Token `json:"tokens"`
}

type Token struct {
	ChainId         uint64 `json:"chain_id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	Logo            string `json:"logo"`
	ContractAddress string `json:"contract_address"`
	ResourceId      string `json:"resource_id"`
}

func GetConfig(c *gin.Context) {
	var res []ChainInfo
	var chains []model.ChainInfo
	db.DB.Model(&model.ChainInfo{}).Where("open=1").Find(&chains)
	for _, chain := range chains {
		var ts []Token
		var tokens []model.TokenInfo
		db.DB.Model(&model.TokenInfo{}).Where("chain_id=?", chain.ChainId).Find(&tokens)
		for _, token := range tokens {
			ts = append(ts, Token{
				ChainId:         uint64(token.ChainId),
				Name:            token.TokenName,
				Symbol:          token.TokenName,
				Decimals:        strconv.FormatInt(token.Decimals, 10),
				Logo:            token.Icon,
				ResourceId:      token.ResourceId,
				ContractAddress: token.TokenAddress,
			})
		}
		res = append(res, ChainInfo{
			ChainId:               uint64(chain.ChainId),
			BlockChainId:          uint64(chain.BlockChainId),
			BridgeContractAddress: chain.BridgeContract,
			Endpoint:              chain.Endpoint,
			ChainName:             chain.ChainName,
			Explorer:              chain.Explorer,
			Logo:                  chain.Logo,
			NativeCoinName:        chain.NativeCoinName,
			NativeCoinDecimals:    chain.NativeCoinDecimals,
			NativeCoinSymbol:      chain.NativeCoinSymbol,
			Tokens:                ts,
		})
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": res,
	})
}
