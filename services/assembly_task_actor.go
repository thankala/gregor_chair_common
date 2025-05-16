package services

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/events"
	"github.com/thankala/gregor_chair_common/interfaces"
)

type AssemblyTaskActor[T any] struct {
	task      enums.Task
	actor     interfaces.AssemblyTask[T]
	server    interfaces.Server
	serverPid *actor.PID
	stopCh    chan struct{}
}

func NewAssemblyTaskActor[T any](actorInstance interfaces.AssemblyTask[T], server interfaces.Server) actor.Producer {
	return func() actor.Receiver {
		return &AssemblyTaskActor[T]{task: actorInstance.Task(), actor: actorInstance, server: server}
	}
}

func (a *AssemblyTaskActor[T]) Receive(ctx *actor.Context) {
	switch event := ctx.Message().(type) {
	case actor.Initialized:
		if value, ok := a.actor.(interfaces.Initializable); ok {
			value.OnInitialized(event, ctx)
		}
	case actor.Started:
		if a.server != nil {
			a.stopCh = make(chan struct{})
			a.serverPid = ctx.SpawnChild(a.server.GetProducer(), "server")
		}
	case actor.Stopped:
		if a.stopCh != nil {
			close(a.stopCh)
		}
	case *events.AssemblyTaskEvent:
		if event.Destination == a.task {
			a.Process(ctx, event)
			return
		}
		if a.server != nil {
			ctx.Send(a.serverPid, event)
		} else {
			ctx.Send(ctx.Parent(), event)
		}
	case *events.OrchestratorEvent:
		if a.server != nil {
			ctx.Send(a.serverPid, event)
		} else {
			ctx.Send(ctx.Parent(), event)
		}
	default:
		panic(fmt.Sprintf("Actor \"%s\" received unknown message: %v", a.task, ctx.Message()))
	}
}

func (a *AssemblyTaskActor[T]) Process(ctx *actor.Context, event *events.AssemblyTaskEvent) {
	if a.task != event.Destination {
		panic(fmt.Sprintf("Actor \"%s\" received message for instance \"%s\"", a.task, event.Source))
	}
	if step, ok := a.actor.Steps()[event.Step]; ok {
		// logger.Get().Info("Processing event", event)
		step(event, ctx)
	} else {
		panic(fmt.Sprintf("Actor \"%s\" received unknown step \"%d\"", a.task, event.Step))
	}
}
