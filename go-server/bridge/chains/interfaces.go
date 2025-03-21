package chains

import (
	"coinstore/bridge/msg"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"math/big"
)

type Connection interface {
	Connect() error
	KeyPrv() *ecdsa.PrivateKey
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	ClientEvm() *ethclient.Client
	ClientTron() *client.GrpcClient
	EnsureHasBytecode(address ethcommon.Address) error
	LatestBlock() (*big.Int, error)
	WaitForBlock(block *big.Int, delay *big.Int) error
	Close()
}

type Router interface {
	Send(message msg.Message) error
}
