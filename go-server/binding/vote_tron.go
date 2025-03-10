package binding

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type VoteTron struct {
	Address string
}

func NewVoteTron(address string) (*VoteTron, error) {
	return &VoteTron{Address: address}, nil
}

func (v *VoteTron) GetProposal(opts *bind.CallOpts, originChainID *big.Int, depositNonce *big.Int, dataHash [32]byte) (IVoteProposal, error) {
	var out IVoteProposal
	var err error

	return out, err

}

func (v *VoteTron) HasVotedOnProposal(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out bool
	var err error

	return out, err
}

func (v *VoteTron) VoteProposal(opts *bind.TransactOpts, originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	var res types.Transaction

	return &res, nil
}

func (v *VoteTron) ExecuteProposal(opts *bind.TransactOpts, originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (*types.Transaction, error) {
	var res types.Transaction

	return &res, nil
}
