package services

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/interfaces"
	"github.com/thankala/gregor_chair_common/messages"
)

type CoordinatorActor[T any] struct {
	coordinator enums.Coordinator
	instance    interfaces.Coordinator[T]
	server      interfaces.Server
	serverPid   *actor.PID
	started     bool
	stopCh      chan struct{}
}

func NewCoordinatorActor[T any](actorInstance interfaces.Coordinator[T], server interfaces.Server) actor.Producer {
	return func() actor.Receiver {
		return &CoordinatorActor[T]{coordinator: actorInstance.Coordinator(), instance: actorInstance, server: server}
	}
}

func (a *CoordinatorActor[T]) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		if value, ok := a.instance.(interfaces.Initializable); ok {
			value.OnInitialized(msg, ctx)
		}
	case actor.Started:
		a.started = true
		if a.server != nil {
			a.stopCh = make(chan struct{})
			a.serverPid = ctx.SpawnChild(a.server.GetProducer(), "server")
		}
	case actor.Stopped:
		a.started = false
		if a.stopCh != nil {
			close(a.stopCh)
		}
	case *messages.AssemblyTaskMessage:
		if a.server != nil {
			ctx.Send(a.serverPid, msg)
		} else {
			ctx.Send(ctx.Parent(), msg)
		}
	case *messages.CoordinatorMessage:
		if msg.Destination == a.coordinator.String() {
			switch msg.Type {
			case enums.PlaceComponent:
				a.instance.PlaceComponent(msg)
			case enums.RequestFixture:
				a.instance.RequestFixture(msg)
			case enums.AttachComponent:
				a.instance.AttachComponent(msg)
			default:
				panic("unhandled default case")
			}
			a.instance.Process(ctx)
		} else {
			if a.server != nil {
				a.server.Send(msg.Source, msg.Destination, msg.Event, msg)
			} else {
				ctx.Send(ctx.Parent(), msg)
			}
		}

		//case *messages.RequestFixture:
		//	//a.instance.RequestFixture(msg)
		//	//a.instance.Process(ctx, msg)
		//case *messages.PlaceComponent:
		//	//a.instance.PlaceComponent(msg)
		//	a.instance.Process(ctx, msg)
		//case *messages.AttachComponent:
		//	//a.instance.AttachComponent(msg)
		//	a.instance.Process(ctx, msg)
	default:
		panic("Unknown message payload")
	}
}