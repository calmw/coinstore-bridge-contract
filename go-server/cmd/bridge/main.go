package main

import (
	"coinstore/bridge"
	"log"
	"os"
)

func main() {
	if err := bridge.Run(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
