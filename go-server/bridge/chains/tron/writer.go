package tron

import (
	"coinstore/binding"
	"coinstore/bridge/config"
	"coinstore/bridge/msg"
	"github.com/calmw/blog"
	"sync"
)

var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4
var WritersTron *WriterTron

type WriterTron struct {
	muVote       *sync.RWMutex
	muExec       *sync.RWMutex
	Cfg          config.Config
	conn         Connection
	voteContract *binding.VoteTron
	log          log15.Logger
	stop         <-chan int
	sysErr       chan<- error // Reports fatal error to core
}

func NewWriterTron(conn Connection, cfg *config.Config, log log15.Logger, stop <-chan int, sysErr chan<- error) *WriterTron {
	voteContract, err := binding.NewVoteTron(cfg.VoteContractAddress)
	if err != nil {
		panic("new vote contract failed")
	}
	writer := WriterTron{
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

func (w *WriterTron) start() error {
	w.log.Debug("Starting Writer...")
	return nil
}

func (w *WriterTron) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	return w.CreateProposal(m)
}
