package main

import (
	"coinstore/bridge"
	"coinstore/bridge/monitor"
	"log"
	"os"
)

func main() {
	go monitor.Start()
	if err := bridge.Run(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
