package tron

import (
	"coinstore/binding"
	"coinstore/bridge/config"
	"coinstore/bridge/msg"
	log "github.com/calmw/clog"
	"sync"
)

var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4
var WritersTron *Writer

type Writer struct {
	muVote       *sync.RWMutex
	muExec       *sync.RWMutex
	Cfg          config.Config
	conn         *Connection
	voteContract *binding.VoteTron
	log          log.Logger
	stop         <-chan int
	sysErr       chan<- error // Reports fatal error to core
}

func NewWriter(conn *Connection, cfg *config.Config, log log.Logger, stop <-chan int, sysErr chan<- error) *Writer {
	voteContract, err := binding.NewVoteTron(cfg.From, cfg.VoteContractAddress, conn.keyStore, conn.keyAccount, conn.connTron)
	if err != nil {
		panic("new vote contract failed")
	}
	writer := Writer{
		muVote:       new(sync.RWMutex),
		muExec:       new(sync.RWMutex),
		Cfg:          *cfg,
		conn:         conn,
		voteContract: voteContract,
		log:          log,
		stop:         stop,
		sysErr:       sysErr,
	}
	WritersTron = &writer
	log.Debug("new writer", "id", cfg.ChainId)
	return &writer
}

func (w *Writer) start() error {
	w.log.Debug("Starting Writer...")
	return nil
}

func (w *Writer) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	return w.CreateProposal(m)
}
