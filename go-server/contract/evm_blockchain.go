package contract

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"reflect"
	"strings"
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

func AbiEncode(opts []interface{}) ([]byte, error) {
	abiStr, err := generateAbi(opts)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(bytes.NewBufferString(abiStr))
	if err != nil {
		return nil, err
	}

	callData, err := parsedABI.Pack("draw", opts...)
	if err != nil {
		return nil, err
	}

	return callData[4:], nil
}

func generateAbi(data []interface{}) (string, error) {
	var inputs string
	for i, d := range data {
		t, err := parseType(d)
		if err != nil {
			return "", err
		}
		inputs += fmt.Sprintf(`{
                "internalType": "%s",
                "name": "parameter%d",
                "type": "%s"
            },`, t, i, t)
	}
	inputs = strings.TrimRight(inputs, ",")

	return fmt.Sprintf(`
[
    {
        "inputs": [
            %s
        ],
        "name": "draw",
        "stateMutability": "pure",
        "type": "function"
    }
]`, inputs), nil

}

func parseType(data interface{}) (string, error) {
	t := reflect.TypeOf(data)
	switch t.String() {
	case "*big.Int":
		return "uint256", nil
	case "[]*big.Int":
		return "uint256[]", nil
	default:
		return "", errors.New("opts type error")
	}
}
