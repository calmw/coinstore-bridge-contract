package binding

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/pkg/errors"
)

type TronAddress struct {
	address string
}

func (tronAddress TronAddress) String() string {
	return tronAddress.address
}

func (tronAddress *TronAddress) Set(s string) error {
	_, err := address.Base58ToAddress(s)
	if err != nil {
		return errors.Wrap(err, "not a valid one address")
	}
	tronAddress.address = s
	return nil
}

func (tronAddress *TronAddress) GetAddress() address.Address {
	addr, err := address.Base58ToAddress(tronAddress.address)
	if err != nil {
		return nil
	}
	return addr
}

func (tronAddress TronAddress) Type() string {
	return "tron-address"
}
