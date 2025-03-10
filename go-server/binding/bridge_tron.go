package binding

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type BridgeTron struct {
	Address string
}

func NewBridgeTron(address string) (*BridgeTron, error) {
	return &BridgeTron{
		Address: address,
	}, nil
}

func (b *BridgeTron) DepositRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	DestinationChainId *big.Int
	Sender             common.Address
	ResourceID         [32]byte
	Data               []byte
}, error) {

	outstruct := new(struct {
		DestinationChainId *big.Int
		Sender             common.Address
		ResourceID         [32]byte
		Data               []byte
	})

	return *outstruct, nil

}
