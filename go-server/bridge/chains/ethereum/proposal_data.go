package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

func ConstructGenericProposalData(metadata []byte) []byte {
	var data []byte

	metadataLen := big.NewInt(int64(len(metadata)))
	fmt.Println(len(math.PaddedBigBytes(metadataLen, 32)), "~~~~~~~~~~~~~~~~~~~_---------------------")
	data = append(data, math.PaddedBigBytes(metadataLen, 32)...) // length of metadata (uint256)
	data = append(data, metadata...)                             // metadata ([]byte)
	return data
}
