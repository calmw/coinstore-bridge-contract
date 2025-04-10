package contract

import (
	"coinstore/utils"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

const (
	ZeroAddress = "0x0000000000000000000000000000000000000000"
	//AdminAccount    = "0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"
	AdminAccount    = "0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9"
	Realyer1Account = "0xD310068976a666D4279F5AdA577DE075e1F32563"
	Realyer2Account = "0xa5b109F231D8b36C8d8fD2b25e99F402FA7e03bE"
	Realyer3Account = "0xC5697941c4fD32391a47db2075802771cBAF34F4"
	AdminRole       = "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	BridgeRole      = "52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f"
	RelayerRole     = "e2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4"
	VoteRole        = "c65b6dc445843af69e7af2fc32667c7d3b98b02602373e2d0a7a047f274806f7"
	ResourceIdUsdt  = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d"
	ResourceIdUsdc  = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1e"
	ResourceIdEth   = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1c"
	ResourceIdWeth  = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1b" // 针对以太坊WETH,需要设置在以太坊和Tantin
)

type ChainConfigs struct {
	BridgeId              int64
	ChainId               int64
	ChainTypeId           int64
	RPC                   string
	BridgeContractAddress string
	VoteContractAddress   string
	TantinContractAddress string
	UsdtAddress           string
	UsdcAddress           string
	EthAddress            string
	WEthAddress           string
	PrivateKey            string
}

var ChainConfig ChainConfigs

func Client(c ChainConfigs) (error, *ethclient.Client) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		log.Fatal("dail failed")
	}
	return nil, client
}

func GetAuth(cli *ethclient.Client) (error, *bind.TransactOpts) {
	privateKeyEcdsa, err := crypto.HexToECDSA(ChainConfig.PrivateKey)

	if err != nil {
		log.Println(err)
		return err, nil
	}
	publicKey := crypto.PubkeyToAddress(privateKeyEcdsa.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return err, nil
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(ChainConfig.ChainId))
	if err != nil {
		log.Println(err)
		return err, nil
	}

	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    auth.Signer,
		Value:     big.NewInt(0),
		GasPrice:  gasPrice,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false,
	}
}

func GetAuthWithValue(cli *ethclient.Client, value *big.Int) (error, *bind.TransactOpts) {
	privateKeyEcdsa, err := crypto.HexToECDSA(ChainConfig.PrivateKey)

	if err != nil {
		log.Println(err)
		return err, nil
	}
	publicKey := crypto.PubkeyToAddress(privateKeyEcdsa.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return err, nil
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(ChainConfig.ChainId))
	if err != nil {
		log.Println(err)
		return err, nil
	}

	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    auth.Signer,
		Value:     value,
		GasPrice:  gasPrice,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false,
	}
}

func InitTantinEnv() {
	coinStoreBridge := os.Getenv("TT_BRIDGE_MAINNET_TEST_DEPLOYER")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:              1,
		ChainId:               202502,
		ChainTypeId:           1,
		RPC:                   "https://rpc.tantin.com",
		BridgeContractAddress: "0x0E55103d7862e4EB499159BDc876A2156A795D06",
		VoteContractAddress:   "0x27B56c6A1C66A78e41A20141e79F8559C33af9b5",
		TantinContractAddress: "0x6f98FEa5fb2A542729Ae667EaC85332d1efbba34",
		UsdtAddress:           "0x2Bf013133aE838B6934B7F96fd43A10EE3FC3e18",
		UsdcAddress:           "0xF3f9629Bf5fC6e40e28444aEA4405dD00e5890AE",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0x99276acEDe57b8dc5632b818DE33B3141DD6FE1d",
		PrivateKey:            privateKeyStr,
	}
}

func InitBscEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:              2,
		ChainId:               56,
		ChainTypeId:           1,
		RPC:                   "https://bsc-mainnet.rpcfast.com?api_key=xbhWBI1Wkguk8SNMu1bvvLurPGLXmgwYeC4S6g2H7WdwFigZSmPWVZRxrskEQwIf",
		BridgeContractAddress: "0x5DbBF032a3ce42D8B6e16c262664Cd98D07802B1",
		VoteContractAddress:   "0xC89b85a50b50f3dD461dbc892e04002d4cd73C54",
		TantinContractAddress: "0x1Ba94cfbe6851309debb6a7C8d3467c214Ada00f",
		UsdtAddress:           "0x55d398326f99059ff775485246999027b3197955",
		UsdcAddress:           "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
		PrivateKey:            privateKeyStr,
	}
}

func InitEthEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:              4,
		ChainId:               1,
		ChainTypeId:           1,
		RPC:                   "https://go.getblock.io/aefd01aa907c4805ba3c00a9e5b48c6b",
		BridgeContractAddress: "0xC0E8a9C9872A6A7E7F5F2999731dec5d798D82B7",
		VoteContractAddress:   "0x6b66eBFA87AaC1dB355B0ec49ECab7F4b32b1b30",
		TantinContractAddress: "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		UsdtAddress:           "0xdac17f958d2ee523a2206206994597c13d831ec7", // Tether USD
		UsdcAddress:           "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // USD Coin
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		PrivateKey:            privateKeyStr,
	}
}
