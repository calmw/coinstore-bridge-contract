package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"math/big"
	"os"
	"strings"
)

func BridgeAdminSetResourceSignatureTron(sigNonce, chainId *big.Int, resourceID [32]byte, assetsType uint8, tokenAddress, tantinAddress common.Address, fee *big.Int, pause bool) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(BridgeSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetResourceSignature",
		sigNonce,
		chainId,
		resourceID,
		assetsType,
		tokenAddress,
		fee,
		pause,
		tantinAddress,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}

func BridgeAdminSetEnvSignatureTron(ks *keystore.KeyStore, ka *keystore.Account, sigNonce *big.Int, voteAddress common.Address, chainId, chainType *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(BridgeSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		chainId,
		voteAddress,
		chainId,
		chainType,
	)
	fmt.Println("~~~~~~~~~ 1 ")
	hash := crypto.Keccak256Hash(parameterBytes[4:])
	signHash, err := ks.SignHash(*ka, hash[:])
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("0x%x", signHash))
	//return GenerateSignatureTron(parameterBytes[4:])
	///

	//coinStoreBridge := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)

	singer := beeCrypto.NewDefaultSigner(privateKey)

	//hash := crypto.Keccak256Hash(parameterBytes[4:])
	// 私钥签名hash
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("0x%x\n", sign)
	return sign, err

	///
}
