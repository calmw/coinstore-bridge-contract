package binding

import (
	"coinstore/bridge/chains/tron/trigger"
	"coinstore/utils"
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type VoteTron struct {
	FromAddress     string
	ContractAddress string
	//keyStore        *keystore.KeyStore
	//keyAccount      *keystore.Account
	cli *client.GrpcClient
}

func NewVoteTron(fromAddress, contractAddress string, cli *client.GrpcClient) (*VoteTron, error) {
	return &VoteTron{
		FromAddress:     fromAddress,
		ContractAddress: contractAddress,
		cli:             cli,
	}, nil
}

func (v *VoteTron) GetProposal(originChainID *big.Int, depositNonce *big.Int, dataHash [32]byte) (IVoteProposal, error) {
	from, _ := utils.TronToEth(v.FromAddress)
	to, _ := utils.TronToEth(v.ContractAddress)
	proposal, err := trigger.GetProposal(from, to, originChainID, depositNonce, dataHash)
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
	from, _ := utils.TronToEth(v.ContractAddress)
	to, _ := utils.TronToEth(v.ContractAddress)
	return trigger.HasVotedOnProposal(from, to, arg0, arg1, arg2)
}

func (v *VoteTron) VoteProposal(originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (string, error) {
	return trigger.VoteProposal(v.cli, v.FromAddress, v.ContractAddress, originChainId, originDepositNonce, resourceId, dataHash)
}

func (v *VoteTron) ExecuteProposal(originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (string, error) {
	return trigger.ExecuteProposal(v.cli, v.FromAddress, v.ContractAddress, originChainId, originDepositNonce, data, resourceId)
}
