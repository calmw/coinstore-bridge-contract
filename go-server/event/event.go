package event

import (
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

type Sig string

// Deposit(
//
//	    uint256 indexed destinationChainId,
//	    bytes32 indexed resourceID,
//	    uint256 indexed depositNonce,
//	    bytes data
//	);

const (
	Deposit Sig = "Deposit(uint256,bytes32,uint256,bytes)"
)

type Event struct {
	EventSignature Sig
	EventName      string
}

var DepositEvent = Event{
	Deposit,
	"DepositEvent",
}

func BuildQuery(contract common.Address, sig Sig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

func (es Sig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

func GetEventName(sig Sig) string {
	aSig := string(sig)
	index := strings.Index(aSig, "(")
	return aSig[:index]
}
