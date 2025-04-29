package signature

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/status-im/keycard-go/hexutils"
	"strings"
)

func SendTransactionUseRlpData(cli *ethclient.Client, rlpDataHex string) (string, error) {
	rawTxBytes := hexutils.HexToBytes(strings.TrimPrefix(rlpDataHex, "0x"))
	tx := new(types.Transaction)
	err := rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return "", err
	}
	err = cli.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}
