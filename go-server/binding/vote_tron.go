package binding

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"math/big"
)

type VoteTron struct {
	Address    string
	keyStore   *keystore.KeyStore
	keyAccount *keystore.Account
	cli        *client.GrpcClient
}

func NewVoteTron(address string, keyStore *keystore.KeyStore, keyAccount *keystore.Account, cli *client.GrpcClient) (*VoteTron, error) {
	return &VoteTron{
		Address:    address,
		keyStore:   keyStore,
		keyAccount: keyAccount,
		cli:        cli,
	}, nil
}

func (v *VoteTron) GetProposal(originChainID *big.Int, depositNonce *big.Int, dataHash [32]byte) (IVoteProposal, error) {
	var out IVoteProposal
	var err error
	///

	//	requestData, err := tron.GenerateBridgeDepositRecordsData(destinationChainId, depositNonce)
	//	if err != nil {
	//		return utils.DepositRecord{}, fmt.Errorf("generateBridgeDepositRecordsData error %v", err)
	//	}
	//	url := fmt.Sprintf("%s/jsonrpc", config.TronApiHost)
	//	if !strings.HasPrefix(from, "0x") {
	//		fromAddress, err := address.Base58ToAddress(from)
	//		if err != nil {
	//			return utils.DepositRecord{}, err
	//		}
	//		from = fromAddress.Hex()
	//	}
	//	if !strings.HasPrefix(to, "0x") {
	//		toAddress, err := address.Base58ToAddress(to)
	//		if err != nil {
	//			return utils.DepositRecord{}, err
	//		}
	//		to = toAddress.Hex()
	//	}
	//	ethCallBody := fmt.Sprintf(`{
	//	"jsonrpc": "2.0",
	//	"method": "eth_call",
	//	"params": [{
	//		"from": "%s",
	//		"to": "%s",
	//		"gas": "0x0",
	//		"gasPrice": "0x0",
	//		"value": "0x0",
	//		"data": "%s"
	//	}, "latest"],
	//	"id": %d
	//}`, from, to, requestData, utils.RandInt(100, 10000))
	//	req, _ := http.NewRequest("POST", url, strings.NewReader(ethCallBody))
	//	req.Header.Add("accept", "application/json")
	//	res, _ := http.DefaultClient.Do(req)
	//	defer res.Body.Close()
	//	body, _ := io.ReadAll(res.Body)
	//	var jsonRpcResponse JsonRpcResponse
	//	err = json.Unmarshal(body, &jsonRpcResponse)
	//	if err != nil {
	//		return utils.DepositRecord{}, errors.New("eth call failed")
	//	}
	//	return utils.ParseBridgeDepositRecordData(hexutils.HexToBytes("197649b0" + strings.TrimPrefix(jsonRpcResponse.Result, "0x")))

	///
	return out, err

}

func (v *VoteTron) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out bool
	var err error

	return out, err
}

func (v *VoteTron) VoteProposal(originChainId *big.Int, originDepositNonce *big.Int, resourceId [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	var res types.Transaction

	return &res, nil
}

func (v *VoteTron) ExecuteProposal(originChainId *big.Int, originDepositNonce *big.Int, data []byte, resourceId [32]byte) (*types.Transaction, error) {
	var res types.Transaction

	return &res, nil
}
