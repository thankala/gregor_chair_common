package interfaces

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
)

type Server interface {
	Send(from string, to string, event enums.Event, msg any)
	Accept(ctx *actor.Context, stopCh <-chan struct{})
	Receive(ctx *actor.Context)
	GetProducer() actor.Producer
}
