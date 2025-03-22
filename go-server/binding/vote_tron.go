package binding

import (
	"coinstore/bridge/tron"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"math/big"
)

type VoteTron struct {
	Address    string
	keyStore   *keystore.KeyStore
	keyAccount *keystore.Account
	cli        *client.GrpcClient
}

func NewVoteTron(address string, keyStore *keystore.KeyStore, keyAccount *keystore.Account, cli *client.GrpcClient) (*VoteTron, error) {
	return &VoteTron{
		Address:    address,
		keyStore:   keyStore,
		keyAccount: keyAccount,
		cli:        cli,
	}, nil
}

func (v *VoteTron) GetProposal(originChainID *big.Int, depositNonce *big.Int, dataHash [32]byte) (IVoteProposal, error) {
	return tron.GetProposal(v.Address, v.Address, originChainID, depositNonce, dataHash)
}

func (v *VoteTron) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out bool
	var err error

	return out, err
}

func (v *VoteTron) VoteProposal(originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	var res types.Transaction

	return &res, nil
}

func (v *VoteTron) ExecuteProposal(originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (*types.Transaction, error) {
	var res types.Transaction

	return &res, nil
}
