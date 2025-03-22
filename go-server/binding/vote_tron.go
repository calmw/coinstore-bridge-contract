package binding

import (
	"coinstore/bridge/tron"
	"github.com/ethereum/go-ethereum/common"
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
	proposal, err := tron.GetProposal(v.Address, v.Address, originChainID, depositNonce, dataHash)
	if err != nil {
		return IVoteProposal{}, err
	}
	return IVoteProposal{
		ResourceId:    proposal.ResourceId,
		DataHash:      proposal.DataHash,
		YesVotes:      proposal.YesVotes,
		NoVotes:       proposal.NoVotes,
		Status:        proposal.Status,
		ProposedBlock: proposal.ProposedBlock,
	}, nil
}

func (v *VoteTron) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return tron.HasVotedOnProposal(v.Address, v.Address, arg0, arg1, arg2)
}

func (v *VoteTron) VoteProposal(originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (string, error) {
	return tron.VoteProposal(v.cli, v.Address, v.Address, v.keyStore, v.keyAccount, originChainId, originDepositNonce, resourceId, dataHash)
}

func (v *VoteTron) ExecuteProposal(originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (string, error) {
	return tron.ExecuteProposal(v.cli, v.Address, v.Address, v.keyStore, v.keyAccount, originChainId, originDepositNonce, data, resourceId)
}
