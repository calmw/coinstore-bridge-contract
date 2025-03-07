// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
	"coinstore/binding/bridge"
	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/chainbridge-utils/msg"
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
	Cfg            Config
	bridgeContract *bridge.Bridge
	log            log15.Logger
}

// NewWriter creates and returns Writer
func NewWriter(cfg *Config, log log15.Logger) *Writer {
	writer := Writer{
		muVote: new(sync.RWMutex),
		muExec: new(sync.RWMutex),
		Cfg:    *cfg,
		log:    log,
	}
	Writers[cfg.chainId] = &writer
	log.Debug("new writer id", "id", cfg.chainId)
	return &writer
}

func (w *Writer) start() error {
	w.log.Debug("Starting Writer...")
	return nil
}

// setContract adds the bound receiver bridgeContract to the Writer
func (w *Writer) setContract(bridge *Bridge.Bridge) {
	w.bridgeContract = bridge
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *Writer) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	switch m.Type {
	case msg.FungibleTransfer:
		return w.createErc20Proposal(m)
	case msg.NonFungibleTransfer:
		return w.createErc721Proposal(m)
	case msg.GenericTransfer:
		return w.CreateGenericDepositProposal(m)
	default:
		w.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}
