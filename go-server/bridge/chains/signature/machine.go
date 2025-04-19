package signature

/// 通过签名机签名
import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
)

func SignAndSendTxEth(cli *ethclient.Client, fromAddress, toAddress common.Address, nonce, gasLimit, chainId uint64, gasPrice, value *big.Int, inputData []byte, apiSecret string) error {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     inputData,
	})
	marshalJSON, err := tx.MarshalJSON()
	if err != nil {
		return err
	}
	fmt.Println("txDataJson")
	fmt.Println(string(marshalJSON))
	taskID := RandInt(100, 10000)
	// 编码数据
	var buf bytes.Buffer
	err = tx.EncodeRLP(&buf)
	txDataRlp := fmt.Sprintf("%x", buf.Bytes())
	fmt.Println("txDataRlp")
	fmt.Println(txDataRlp)
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chainId,
		strings.ToLower(fromAddress.String()),
		taskID,
		txDataRlp,
		apiSecret,
	)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: fromAddress.String(),
		TxData:      txDataRlp,
		TaskID:      taskID,
		ChainID:     int(chainId),
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}

	res, err := GetSignedRlpData("https://10.234.99.69:8088/signature/sign", postData)
	fmt.Println("RequestWithPem error:", err)
	var machineResp MachineResp
	err = json.Unmarshal(res, &machineResp)
	if err != nil {
		return err
	}
	if machineResp.Code != 200 {
		return errors.New("signature machine error")
	}
	// 发送交易
	txHash, err := SendTransactionUseRlpData(cli, machineResp.Data)
	fmt.Println("SendTransactionFromRlpData:")
	fmt.Println(err, txHash)
	return err
}

// SignAndSendTxTron  728126428
func SignAndSendTxTron(chainId int, fromAddress string, UnsignedRawData []byte, apiSecret string) ([]byte, error) {

	taskID := RandInt(100, 10000)
	// 编码数据
	txDataRlp := fmt.Sprintf("%x", UnsignedRawData)
	fmt.Println("txDataRlp")
	fmt.Println(txDataRlp)
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chainId,
		strings.ToLower(fromAddress),
		taskID,
		txDataRlp,
		apiSecret,
	)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: strings.ToLower(fromAddress),
		TxData:      txDataRlp,
		TaskID:      taskID,
		ChainID:     chainId,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}

	res, err := GetSignedRlpData("https://10.234.99.69:8088/signature/sign", postData)
	fmt.Println("RequestWithPem error:", err)
	var machineResp MachineResp
	err = json.Unmarshal(res, &machineResp)
	if err != nil {
		return nil, err
	}
	if machineResp.Code != 200 {
		return nil, errors.New("signature machine error")
	}

	return hexutils.HexToBytes(machineResp.Data), nil
}
