package abi

import (
	"coinstore/binding"
	"coinstore/utils"
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
	fmt.Println(fmt.Sprintf("0x%x", parameterBytes[4:]))
	fmt.Println("~~~~~~~~~ 2 ")
	fmt.Println(binding.OwnerAccount)
	fmt.Println(utils.TronToEth(binding.OwnerAccount))
	fmt.Println(utils.TronToEth("TDtRnGZz38MQQ3k1K2QiyoajQSbAhAVctL"))
	fmt.Println(utils.EthToTron("0x8840E6C55B9ADA326D211D818C34A994AECED808"))
	fmt.Println(utils.EthToTron("0x5f28b0F238917aC64805dD5bc47A043e16926AC4"))
	fmt.Println(utils.EthToTron("0x2aF91752A2379a2Cf60C3a5c929a1E534185c85C"))
	fmt.Println(utils.EthToTron("0xC879c7dEfd8B4aF6d6DfFc33b73A4aAc12a461C9"))
	hash := crypto.Keccak256Hash(parameterBytes[4:])
	signHash, err := ks.SignHash(*ka, hash[:])
	if err != nil {
		return nil, err
	}
	fmt.Println("~~~~~~~~~ 3 ")
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
	fmt.Println("~~~~~~~~~ 4 ")
	fmt.Printf("0x%x\n", sign)
	return sign, err

	///
}
