package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"os"
)

var TantinSig = `[{
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "tokenAddress",
		"type": "address"
	}, {
		"internalType": "uint256",
		"name": "amount",
		"type": "uint256"
	}],
	"name": "checkAdminWithdrawSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	}, {
		"internalType": "bytes32",
		"name": "resourceID",
		"type": "bytes32"
	}, {
		"internalType": "enum ITantinBridge.AssetsType",
		"name": "assetsType",
		"type": "uint8"
	}, {
		"internalType": "address",
		"name": "tokenAddress",
		"type": "address"
	}, {
		"internalType": "bool",
		"name": "burnable",
		"type": "bool"
	}, {
		"internalType": "bool",
		"name": "mintable",
		"type": "bool"
	}, {
		"internalType": "bool",
		"name": "pause",
		"type": "bool"
	}],
	"name": "checkAdminSetTokenSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "user",
		"type": "address"
	}],
	"name": "checkAdminRemoveBlacklistSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "user",
		"type": "address"
	}],
	"name": "checkAdminAddBlacklistSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "bridgeAddress",
		"type": "address"
	}],
	"name": "checkAdminSetEnvSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "address",
		"name": "recipient",
		"type": "address"
	}],
	"name": "checkDepositSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "pure",
	"type": "function"
}]`

var BridgeSig = `[{
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	},{
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	},{
		"internalType": "bytes32",
		"name": "resourceID",
		"type": "bytes32"
	}, {
		"internalType": "enum IBridge.AssetsType",
		"name": "assetsType",
		"type": "uint8"
	}, {
		"internalType": "address",
		"name": "tokenAddress",
		"type": "address"
	}, {
		"internalType": "uint256",
		"name": "fee",
		"type": "uint256"
	}, {
		"internalType": "bool",
		"name": "pause",
		"type": "bool"
	}, {
		"internalType": "address",
		"name": "tantinAddress",
		"type": "address"
	}],
	"name": "checkAdminSetResourceSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	},{
		"internalType": "address",
		"name": "voteAddress_",
		"type": "address"
	}, {
		"internalType": "uint256",
		"name": "chainId_",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainType_",
		"type": "uint256"
	}],
	"name": "checkAdminSetEnvSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}]`

var VoteSig = `[{
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "relayerAddress",
		"type": "address"
	}],
	"name": "checkAdminRemoveRelayerSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "originChainID",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "originDepositNonce",
		"type": "uint256"
	}, {
		"internalType": "bytes32",
		"name": "dataHash",
		"type": "bytes32"
	}],
	"name": "cancelProposal",
	"outputs": [],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "relayerAddress",
		"type": "address"
	}],
	"name": "checkAdminAddRelayerSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "newThreshold",
		"type": "uint256"
	}],
	"name": "checkAdminChangeRelayerThresholdSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}, {
	"inputs": [{
		"internalType": "uint256",
		"name": "sigNonce",
		"type": "uint256"
	}, {
		"internalType": "address",
		"name": "bridgeAddress_",
		"type": "address"
	}, {
		"internalType": "address",
		"name": "tantinAddress_",
		"type": "address"
	}, {
		"internalType": "uint256",
		"name": "expiry_",
		"type": "uint256"
	}, {
		"internalType": "uint256",
		"name": "relayerThreshold_",
		"type": "uint256"
	}],
	"name": "checkAdminSetEnvSignature",
	"outputs": [{
		"internalType": "bool",
		"name": "",
		"type": "bool"
	}],
	"stateMutability": "nonpayable",
	"type": "function"
}]`

func GenerateSignature(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("TT_BRIDGE_SIGN")
	//coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	//privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	fmt.Println(privateKeyStr)
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
