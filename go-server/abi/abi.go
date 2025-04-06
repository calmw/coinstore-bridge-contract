package abi

import (
	"coinstore/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"os"
)

var TantinSig = `[
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "tokenAddress",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "checkAdminWithdrawSignature",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "resourceID",
          "type": "bytes32"
        },
        {
          "internalType": "enum ITantinBridge.AssetsType",
          "name": "assetsType",
          "type": "uint8"
        },
        {
          "internalType": "address",
          "name": "tokenAddress",
          "type": "address"
        },
        {
          "internalType": "bool",
          "name": "burnable",
          "type": "bool"
        },
        {
          "internalType": "bool",
          "name": "mintable",
          "type": "bool"
        },
        {
          "internalType": "bool",
          "name": "pause",
          "type": "bool"
        }
      ],
      "name": "checkAdminSetTokenSignature",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        }
      ],
      "name": "checkAdminRemoveBlacklistSignature",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        }
      ],
      "name": "checkAdminAddBlacklistSignature",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "bridgeAddress",
          "type": "address"
        }
      ],
      "name": "checkAdminSetEnvSignature",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        }
      ],
      "name": "checkDepositSignature",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    }
]`

var BridgeSig = `[
	{
		"inputs": [{
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      }],
		"name": "recipientSignature",
		"outputs": [],
		"stateMutability": "view",
		"type": "function"
	}
]`

var VoteSig = `[
	{
		"inputs": [{
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      }],
		"name": "recipientSignature",
		"outputs": [],
		"stateMutability": "view",
		"type": "function"
	}
]`

func GenerateSignature(parameter []byte) ([]byte, error) {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}
	singer := beeCrypto.NewDefaultSigner(privateKey)

	hash := crypto.Keccak256Hash(parameter)
	// 私钥签名hash
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("0x%x\n", sign)
	return sign, err
}
