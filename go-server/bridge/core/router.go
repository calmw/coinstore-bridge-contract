package core

import (
	"coinstore/bridge/msg"
	"fmt"
	log "github.com/calmw/blog"
	"sync"
)

type Writer interface {
	ResolveMessage(message msg.Message) bool
}

type Router struct {
	registry map[msg.ChainId]Writer
	lock     *sync.RWMutex
	log      log.Logger
}

func NewRouter(log log.Logger) *Router {
	return &Router{
		registry: make(map[msg.ChainId]Writer),
		lock:     &sync.RWMutex{},
		log:      log,
	}
}

func (r *Router) Send(msg msg.Message) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.log.Trace("Routing message", "src", msg.Source, "dest", msg.Destination, "nonce", msg.DepositNonce, "rId", msg.ResourceId.Hex())
	w := r.registry[msg.Destination]
	if w == nil {
		return fmt.Errorf("unknown destination chainId: %d", msg.Destination)
	}

	go w.ResolveMessage(msg)
	return nil
}

func (r *Router) Listen(id msg.ChainId, w Writer) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.log.Debug("Registering new chain in interface", "id", id)
	r.registry[id] = w
}
