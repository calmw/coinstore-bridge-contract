package signature

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
	"strings"
)

func GenerateVoteProposalInputData(originChainId, originDepositNonce *big.Int, resourceId, dataHash [32]byte) ([]byte, error) {
	abiJson := `[{
    "inputs": [
      {
        "internalType": "uint256",
        "name": "originChainId",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "originDepositNonce",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "resourceId",
        "type": "bytes32"
      },
      {
        "internalType": "bytes32",
        "name": "dataHash",
        "type": "bytes32"
      }
    ],
    "name": "voteProposal",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }]`
	contractAbi, _ := abi.JSON(strings.NewReader(abiJson))
	return contractAbi.Pack("voteProposal",
		originChainId,
		originDepositNonce,
		resourceId,
		dataHash,
	)
}

func GenerateExecuteProposalInputData(originChainId, originDepositNonce *big.Int, data []byte) ([]byte, error) {
	abiJson := `[{
    "inputs": [
      {
        "internalType": "uint256",
        "name": "originChainId",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "originDepositNonce",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "name": "executeProposal",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }]`
	contractAbi, _ := abi.JSON(strings.NewReader(abiJson))
	return contractAbi.Pack("voteProposal",
		originChainId,
		originDepositNonce,
		data,
	)
}
