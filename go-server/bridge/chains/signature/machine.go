package signature

/// 通过签名机签名
import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strings"
)

func SignAndSendTxEth(cli *ethclient.Client, chainId uint64, tx *types.Transaction, apiSecret string) error {
	marshalJSON, err := tx.MarshalJSON()
	if err != nil {
		return err
	}
	sigAccount := os.Getenv("SIG_ACCOUNT_EVM")
	if len(sigAccount) <= 0 {
		sigAccount = "0x1933ccd14cafe561d862e5f35d0de75322a55412"
	}
	fromAddress := common.HexToAddress(sigAccount)
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
	fmt.Println("1")
	fmt.Println(sigStr)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: fromAddress.String(),
		TxData:      txDataRlp,
		TaskID:      taskID,
		ChainID:     int(chainId),
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}

	res, err := GetSignedRlpData("https://18.141.210.154:8088/signature/sign", postData)
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
func SignAndSendTxTron(chainId int, UnsignedRawData []byte, apiSecret string) ([]byte, error) {
	sigAccount := os.Getenv("SIG_ACCOUNT_TRON")
	if len(sigAccount) <= 0 {
		sigAccount = "TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n"
	}
	chainId = 728126428
	taskID := RandInt(100, 10000)
	// 签名数据
	h256h := sha256.New()
	h256h.Write(UnsignedRawData)
	hash := h256h.Sum(nil)
	rawData := fmt.Sprintf("%x", hash)

	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chainId,
		strings.ToLower(sigAccount),
		taskID,
		rawData,
		apiSecret,
	)
	fmt.Println("2")
	fmt.Println(sigStr)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: sigAccount,
		TxData:      rawData,
		TaskID:      taskID,
		ChainID:     chainId,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}

	res, err := GetSignedRlpData("https://18.141.210.154:8088/signature/sign", postData)
	var machineResp MachineResp
	err = json.Unmarshal(res, &machineResp)
	if err != nil {
		return nil, err
	}
	if machineResp.Code != 200 {
		return nil, fmt.Errorf("signature machine error: %v", machineResp.Message)
	}
	sigBytes, err := hex.DecodeString(machineResp.Data)
	if err != nil {
		return nil, err
	}

	return sigBytes, nil
}
