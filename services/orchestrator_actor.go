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
}

func NewOrchestratorActor[T any](actorInstance interfaces.Orchestrator[T], server interfaces.Server) actor.Producer {
	return func() actor.Receiver {
		return &OrchestratorActor[T]{orchestrator: actorInstance.Orchestrator(), instance: actorInstance, server: server}
	}
}

func (a *OrchestratorActor[T]) Receive(ctx *actor.Context) {
	switch event := ctx.Message().(type) {
	case actor.Initialized:
		if value, ok := a.instance.(interfaces.Initializable); ok {
			value.OnInitialized(event, ctx)
		}
	case actor.Started:
		a.started = true
		if a.server != nil {
			a.serverPid = ctx.SpawnChild(a.server.GetProducer(), "server")
		}
	case actor.Stopped:
		a.started = false
	case *events.AssemblyTaskEvent:
		if a.server != nil {
			ctx.Send(a.serverPid, event)
		} else {
			ctx.Send(ctx.Parent(), event)
		}
	case *events.OrchestratorEvent:
		if event.Type == enums.AssemblyStarted {
			a.instance.StartAssembly(ctx, event)
			return
		}
		if event.Destination == a.orchestrator {
			a.instance.Process(ctx, event)
		} else {
			if a.server != nil {
				ctx.Send(a.serverPid, event)
			} else {
				ctx.Send(ctx.Parent(), event)
			}
		}
	default:
		panic("Unknown message payload")
	}
}
