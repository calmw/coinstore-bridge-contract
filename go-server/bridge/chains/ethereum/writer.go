// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"coinstore/binding"
	"coinstore/bridge/config"
	"coinstore/bridge/msg"
	"github.com/ChainSafe/log15"
	"sync"
)

var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4
var Writers = map[int]*Writer{}

type Writer struct {
	muVote         *sync.RWMutex
	muExec         *sync.RWMutex
	Cfg            config.Config
	conn           Connection
	bridgeContract *binding.Bridge
	log            log15.Logger
	stop           <-chan int
	sysErr         chan<- error // Reports fatal error to core
}

func NewWriter(conn Connection, cfg *config.Config, log log15.Logger, stop <-chan int, sysErr chan<- error) *Writer {
	writer := Writer{
		muVote: new(sync.RWMutex),
		muExec: new(sync.RWMutex),
		Cfg:    *cfg,
		conn:   conn,
		log:    log,
		stop:   stop,
		sysErr: sysErr,
	}
	Writers[cfg.ChainId] = &writer
	log.Debug("new writer id", "id", cfg.ChainId)
	return &writer
}

func (w *Writer) start() error {
	w.log.Debug("Starting Writer...")
	return nil
}

func (w *Writer) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	//switch m.Type {
	//case msg.FungibleTransfer:
	//	return w.createErc20Proposal(m)
	//case msg.NonFungibleTransfer:
	//	return w.createErc721Proposal(m)
	//case msg.GenericTransfer:
	//	return w.CreateGenericDepositProposal(m)
	//default:
	//	w.log.Error("Unknown message type received", "type", m.Type)
	//	return false
	//}

	return false
}
