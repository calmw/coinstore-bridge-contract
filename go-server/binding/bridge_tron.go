package binding

type BridgeTron struct {
	Address string
}

func NewBridgeTron(address string) (*BridgeTron, error) {
	return &BridgeTron{
		Address: address,
	}, nil
}
