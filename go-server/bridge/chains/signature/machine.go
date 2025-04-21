package signature

/// 通过签名机签名
import (
	"bytes"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
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

func SignAndSendTxEth(cli *ethclient.Client, fromAddress common.Address, chainId uint64, tx *types.Transaction, apiSecret string) error {
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
	// 签名数据
	rawData := fmt.Sprintf("%x", UnsignedRawData)
	fmt.Println("rawData")
	fmt.Println(rawData)
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chainId,
		strings.ToLower(fromAddress),
		taskID,
		rawData,
		apiSecret,
	)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: fromAddress,
		TxData:      rawData,
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
		return nil, fmt.Errorf("signature machine error: %v", machineResp.Message)
	}

	fmt.Println("signature:")
	fmt.Println(machineResp.Data)
	sigBytes1, _ := hex.DecodeString(machineResp.Data)
	sigBytes2 := hexutils.HexToBytes(machineResp.Data)

	fmt.Println(len(sigBytes1), "@@@@@@@@@@11111~~~~~~~~~~~~")
	fmt.Println(len(sigBytes2), "@@@@@@@@@@11111~~~~~~~~~~~~")
	// 示例使用
	//derSignature := "3046022100b488d449a2a9244c164301455ffc8b6f5cdb5881ac4c96b6df94dc153e06ddb7022100cb0e1a5eedf1db07268dc637f39d04bf0c311625930f01167076b7d0e94b90d3"

	rawSignature, err := DerToRawSignature(machineResp.Data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, fmt.Errorf("error: %v", err)
	}

	sigBytes, err := hex.DecodeString(rawSignature)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(sigBytes), "1122@~~~~~~~~~~~~")

	return sigBytes, nil
}

// DER签名结构
type ECDSASignature struct {
	R, S *big.Int
}

// DER转原始签名
func DerToRawSignature(derSignature string) (string, error) {
	// 1. 解码十六进制DER签名
	derBytes, err := hex.DecodeString(derSignature)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex: %v", err)
	}

	// 2. 解析DER序列
	var sig ECDSASignature
	_, err = asn1.Unmarshal(derBytes, &sig)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal DER: %v", err)
	}

	// 3. 将r和s转换为32字节的数组
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	// 4. 确保r和s都是32字节
	r32 := make([]byte, 32)
	s32 := make([]byte, 32)

	// 从后往前复制，确保对齐
	copy(r32[32-len(rBytes):], rBytes)
	copy(s32[32-len(sBytes):], sBytes)

	// 5. 合并r和s
	rawSignature := make([]byte, 64)
	copy(rawSignature[:32], r32)
	copy(rawSignature[32:], s32)

	V := make([]byte, 1)
	if sig.R.Uint64()%2 == 0 {
		V = big.NewInt(27).Bytes()
	} else {
		V = big.NewInt(28).Bytes()
	}
	var res []byte
	res = append(res, rawSignature...)
	res = append(res, V...)
	// 6. 返回十六进制字符串
	return hex.EncodeToString(res), nil
}
