package model

import (
	"coinstore/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ChainInfo struct {
	gorm.Model
	ChainName          string          `gorm:"chain_name;comment:'链名称'" json:"chain_name"`
	Logo               string          `gorm:"chain_name;type:longtext;comment:'链logo'" json:"logo"`
	Explorer           string          `gorm:"chain_name;type:longtext;comment:'浏览器地址'" json:"explorer"`
	NativeCoinName     string          `gorm:"chain_name;type:longtext;comment:'native coin'" json:"native_coin_name"`
	NativeCoinSymbol   string          `gorm:"chain_name;type:longtext;comment:'native coin'" json:"native_coin_symbol"`
	NativeCoinDecimals string          `gorm:"chain_name;type:longtext;comment:'native coin'" json:"native_coin_decimals"`
	ChainId            int             `gorm:"chain_id;comment:'自定义链ID'" json:"chain_id"`
	BlockChainId       int             `gorm:"chain_id;comment:'链ID'" json:"block_chain_id"`
	ChainType          int             `gorm:"chain_id;comment:'自定义链类型ID,1 EVM 2 TRON'" json:"-"`
	Endpoint           string          `gorm:"endpoint;comment:'RPC'" json:"-"`
	From               string          `gorm:"from;comment:'签名账户地址'" json:"-"`
	PrivateKey         string          `gorm:"private_key;comment:'签名账户私钥'" json:"-"`
	FreshStart         bool            `gorm:"fresh_start;comment:'启动时，是否从设置的起始块高开始.true 从设置的块高开始，false 从block store开始'" json:"-"`
	LatestBlock        bool            `gorm:"latest_block;comment:'启动时，是否从设置的起始块高开始.true 从区块链当前的块高开始'" json:"-"`
	BridgeContract     string          `gorm:"bridge_contract;comment:'bridge合约地址'" json:"-"`
	VoteContract       string          `gorm:"vote_contract;comment:'vote合约地址'" json:"-"`
	TantinContract     string          `gorm:"tantin_contract;comment:'tantin合约地址'" json:"-"`
	GasLimit           int64           `gorm:"gas_limit;comment:''" json:"-"`
	MaxGasPrice        int64           `gorm:"max_gas_price;comment:''" json:"-"`
	MinGasPrice        int64           `gorm:"min_gas_price;comment:''" json:"-"`
	Http               int             `gorm:"http;default:1;comment:'http or ws. 1 http 0 ws'" json:"-"`
	StartBlock         decimal.Decimal `gorm:"type:bigint(30);start_block;comment:'起始块高'" json:"-"`
	BlockStore         decimal.Decimal `gorm:"type:bigint(30);block_store;comment:'扫过的块高'" json:"-"`
	BlockConfirmations int64           `gorm:"block_confirmations;comment:'待确认块高差'" json:"-"`
	Open               int64           `gorm:"open;comment:'是否开启 1 开启 0关闭'；default:1" json:"open"`
}

func GetConfigByChainId(tx *gorm.DB, chainId int) (ChainInfo, error) {
	var cfg ChainInfo
	err := tx.Model(&ChainInfo{}).Where("chain_id=?", chainId).First(&cfg).Error
	if err != nil {
		return ChainInfo{}, err
	}
	return cfg, nil
}

func GetAllConfig() ([]ChainInfo, error) {
	var cfgs []ChainInfo
	err := db.DB.Model(&ChainInfo{}).Where("open>0").Find(&cfgs).Error
	return cfgs, err
}
