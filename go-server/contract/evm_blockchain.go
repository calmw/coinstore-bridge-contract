package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

const (
	ZeroAddress    = "0x0000000000000000000000000000000000000000"
	AdminRole      = "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	BridgeRole     = "52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f"
	RelayerRole    = "e2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4"
	VoteRole       = "c65b6dc445843af69e7af2fc32667c7d3b98b02602373e2d0a7a047f274806f7"
	ResourceIdUsdt = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d"
	ResourceIdUsdc = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1e"
	ResourceIdEth  = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1c"
	ResourceIdWeth = "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1b" // 针对以太坊WETH,需要设置在以太坊和Tantin
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
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("TT_BRIDGE_SIGN")
	}
	ChainConfig = ChainConfigs{
		BridgeId:    1,
		ChainId:     12302,
		ChainTypeId: 1,
		RPC:         "https://testrpc.tantin.com",
		//RPC:                   "https://testrpcdex.tantin.com",
		BridgeContractAddress: "0x36A455C2e103bfd475e448e6Ff9ea597c4468ee5",
		VoteContractAddress:   "0xD9A7D83aeb5CA79196DdC9a87f274DBF6cf82d21",
		TantinContractAddress: "0x049dB031F482B3EFEA6649277f3441eaDb3B7085",
		UsdtAddress:           "0x53F1BAA532710FC1FEE8a66433bE6c6fE823fCE9",
		UsdcAddress:           "0x87386337645860720009341caD33C6652806aF6f",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0x167BE2dDDcEc733B5d9f0CE1256D2f0D8EeC7058",
		PrivateKey:            privateKeyStr,
	}
}

func InitTantinEnvProd() {
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("TT_BRIDGE_SIGN")
	}
	ChainConfig = ChainConfigs{
		BridgeId:              1,
		ChainId:               12301,
		ChainTypeId:           1,
		RPC:                   "https://rpc.tantin.com",
		BridgeContractAddress: "0x0AAa93530Af315d910DFb3A745D010Bf2d110b4e",
		VoteContractAddress:   "0x579eB412D19f3Bcd628c5e38e92Ac2B9B16ACC65",
		TantinContractAddress: "0xdbADa77D545Ca263Df8dDf9c8bf8109383f6e50a",
		UsdtAddress:           "0x2Bf013133aE838B6934B7F96fd43A10EE3FC3e18",
		UsdcAddress:           "0xF3f9629Bf5fC6e40e28444aEA4405dD00e5890AE",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0x99276acEDe57b8dc5632b818DE33B3141DD6FE1d",
		PrivateKey:            privateKeyStr,
	}
}

func InitBscEnv() {
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("TT_BRIDGE_SIGN")
	}
	ChainConfig = ChainConfigs{
		BridgeId:              2,
		ChainId:               97,
		ChainTypeId:           1,
		RPC:                   "https://data-seed-prebsc-2-s3.bnbchain.org:8545",
		BridgeContractAddress: "0x5389d4eae713F7b7F57e88faaaCF5E7aEe06E0ad",
		VoteContractAddress:   "0xfa363746c7a69ea50C7019e789Ab96e55E5D3aD5",
		TantinContractAddress: "0xA84BDAf841AD6980b205df56F170df279698845a",
		UsdtAddress:           "0x4b62Da623b5aAfE4BAEe909e1fBB321b96887B3D",
		UsdcAddress:           "0xA94706880640B461E25034277E3f8d625B730Bc6",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0x43f66dB67821e38BF935924c999B94dBD24Bd35f",
		PrivateKey:            privateKeyStr,
	}
}

func InitBscProdEnv() {
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("TT_BRIDGE_SIGN")
	}
	ChainConfig = ChainConfigs{
		BridgeId:              2,
		ChainId:               56,
		ChainTypeId:           1,
		RPC:                   "https://bsc-mainnet.infura.io/v3/59ec080dc74d4af893ea04bfe2b168b5",
		BridgeContractAddress: "0x221441b7EEae32BFdA24a7d2de71748E20966A49",
		VoteContractAddress:   "0xEEBBae1445AD3c1a7d98F51AE12E8D1132553f13",
		TantinContractAddress: "0x88bbadEFABcd8E2bEF70c85b8985b464dceD4187",
		UsdtAddress:           "0x55d398326f99059ff775485246999027b3197955",
		UsdcAddress:           "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
		PrivateKey:            privateKeyStr,
	}
}

func InitEthEnv() {
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("TT_BRIDGE_SIGN")
	}
	ChainConfig = ChainConfigs{
		BridgeId:    4,
		ChainId:     11155111,
		ChainTypeId: 1,
		RPC:         "https://sepolia.drpc.org",
		//RPC: "https://sepolia.infura.io",
		//RPC:                   "https://ethereum-sepolia-rpc.publicnode.com",
		BridgeContractAddress: "0xEf495bdab97F796091346Db4A46Ca0E82bFeA31f",
		VoteContractAddress:   "0x6DdF0AB701Dff9f40bB3D09Ba67D43C1d3890d1E",
		TantinContractAddress: "0x9c06c39F095d85591fc4dd33E33C9f32D3669233",
		UsdtAddress:           "0x94Bbc0cc03245Ec1f9B5d7134fB3A9D579ADc3c9", // Tether USD
		UsdcAddress:           "0x43f66dB67821e38BF935924c999B94dBD24Bd35f", // USD Coin
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0xA94706880640B461E25034277E3f8d625B730Bc6",
		PrivateKey:            privateKeyStr,
	}
}

func InitEthEnvProd() {
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("TT_BRIDGE_SIGN")
	}
	ChainConfig = ChainConfigs{
		BridgeId:              4,
		ChainId:               1,
		ChainTypeId:           1,
		RPC:                   "https://mainnet.infura.io/v3/59ec080dc74d4af893ea04bfe2b168b5",
		BridgeContractAddress: "0x221441b7EEae32BFdA24a7d2de71748E20966A49",
		VoteContractAddress:   "0xEEBBae1445AD3c1a7d98F51AE12E8D1132553f13",
		TantinContractAddress: "0x88bbadEFABcd8E2bEF70c85b8985b464dceD4187",
		UsdtAddress:           "0xdac17f958d2ee523a2206206994597c13d831ec7", // Tether USD
		UsdcAddress:           "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // USD Coin
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		PrivateKey:            privateKeyStr,
	}
}
