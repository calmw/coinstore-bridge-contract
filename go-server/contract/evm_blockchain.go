package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
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
