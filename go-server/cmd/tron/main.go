package main

import (
	"coinstore/blockchain"
	"coinstore/services"
	"context"
	"fmt"
)

func main() {
	services.InitTronEnv()

	err, cli := blockchain.Client(blockchain.ChainConfig)

	fmt.Println(err, cli)
	fmt.Println(cli.ChainID(context.Background()))
}
