package services

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/events"
	"github.com/thankala/gregor_chair_common/interfaces"
)

type OrchestratorActor[T any] struct {
	orchestrator enums.Task
	instance     interfaces.Orchestrator[T]
	server       interfaces.Server
	serverPid    *actor.PID
	started      bool
	stopCh       chan struct{}
}

func NewOrchestratorActor[T any](actorInstance interfaces.Orchestrator[T], server interfaces.Server) actor.Producer {
	return func() actor.Receiver {
		return &OrchestratorActor[T]{orchestrator: actorInstance.Orchestrator(), instance: actorInstance, server: server}
	}
}

func (a *OrchestratorActor[T]) Receive(ctx *actor.Context) {
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
	case *events.AssemblyTaskEvent:
		// time.Sleep(time.Millisecond * 10)
		if a.server != nil {
			ctx.Send(a.serverPid, msg)
		} else {
			ctx.Send(ctx.Parent(), msg)
		}
	case *events.OrchestratorEvent:
		// time.Sleep(time.Millisecond * 10)
		if msg.Destination == a.orchestrator {
			switch msg.Type {
			case enums.ComponentPlaced:
				a.instance.ComponentPlaced(msg)
			case enums.FixtureRequested:
				a.instance.FixtureRequested(msg)
			case enums.ComponentAttached:
				a.instance.ComponentAttached(msg)
			default:
				panic("unhandled default case")
			}
			a.instance.Process(ctx)
		} else {
			if a.server != nil {
				ctx.Send(a.serverPid, msg)
			} else {
				ctx.Send(ctx.Parent(), msg)
			}
		}
	default:
		panic("Unknown message payload")
	}
}
