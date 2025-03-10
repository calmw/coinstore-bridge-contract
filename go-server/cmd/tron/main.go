package main

import (
	"coinstore/contract"
	"coinstore/services"
	"context"
	"fmt"
)

func main() {
	services.InitTronEnv()

	err, cli := contract.Client(contract.ChainConfig)

	fmt.Println(err, cli)
	fmt.Println(cli.ChainID(context.Background()))
}
