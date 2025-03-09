package chains

import (
	"coinstore/bridge/msg"
)

type Router interface {
	Send(message msg.Message) error
}
