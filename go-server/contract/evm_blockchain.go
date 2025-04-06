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
	AdminAccount   = "0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"
	AdminRole      = "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	VoteRole       = "c65b6dc445843af69e7af2fc32667c7d3b98b02602373e2d0a7a047f274806f7"
	ResourceIdUsdt = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d"
	ResourceIdCoin = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1c"
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
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:              1,
		ChainId:               202502,
		ChainTypeId:           1,
		RPC:                   "https://rpc.tantin.com",
		BridgeContractAddress: "0x2395DDA69077620d44C9b974d028421Eaf802037",
		VoteContractAddress:   "0xbe1e484aBC7bd4e7F2e81eFAd784271790885e1B",
		TantinContractAddress: "0x6944B2479f745C03c8796D4DCd90c15b5155819F",
		UsdtAddress:           "0x2Bf013133aE838B6934B7F96fd43A10EE3FC3e18",
		UsdcAddress:           "0xe82F64b73D58C85803D78be00407fD44c6DeBe63",
		PrivateKey:            privateKeyStr,
	}
}

func InitBscEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:              2,
		ChainId:               5611,
		ChainTypeId:           1,
		RPC:                   "https://opbnb-testnet-rpc.bnbchain.org",
		BridgeContractAddress: "0xca540876A5c64eB1A0E51115CF7a5b2687F6e0d2",
		VoteContractAddress:   "0x7EC7dca61c29773466D33aEB9e4f7adbBA960Ca1",
		TantinContractAddress: "0x09125BB80eb099073b392637De2b6f3A42f7D1aC",
		UsdtAddress:           "0x671b21826BdFB241aCa2Dd49dD6C0B96A9309455",
		UsdcAddress:           "0x740b6892bFe90D8b5B926782761e5F9F3eaCC1A1",
		PrivateKey:            privateKeyStr,
	}
}

func InitEthEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:    4,
		ChainId:     11155111,
		ChainTypeId: 1,
		//RPC:         "https://sepolia.infura.io/v3/732f6502b35c486fb07e333b32e89c04",
		RPC:                   "https://sepolia.drpc.org",
		BridgeContractAddress: "0xC0E8a9C9872A6A7E7F5F2999731dec5d798D82B7",
		VoteContractAddress:   "0x6b66eBFA87AaC1dB355B0ec49ECab7F4b32b1b30",
		TantinContractAddress: "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		UsdtAddress:           "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b", // Tether USD
		UsdcAddress:           "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", // USD Coin
		PrivateKey:            privateKeyStr,
	}
}
