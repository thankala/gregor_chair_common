package services

import (
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/interfaces"
	"github.com/thankala/gregor_chair_common/messages"
)

type AssemblyTaskActor[T any] struct {
	task      enums.AssemblyTask
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
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		if value, ok := a.actor.(interfaces.Initializable); ok {
			value.OnInitialized(msg, ctx)
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
	case *messages.AssemblyTaskMessage:
		if msg.Task == a.task {
			a.Process(ctx, msg)
			return
		}
		if a.server != nil {
			ctx.Send(a.serverPid, msg)
		} else {
			ctx.Send(ctx.Parent(), msg)
		}
	case *messages.CoordinatorMessage:
		if a.server != nil {
			ctx.Send(a.serverPid, msg)
		} else {
			ctx.Send(ctx.Parent(), msg)
		}
	default:
		panic(fmt.Sprintf("Actor \"%s\" received unknown message: %v", a.task, ctx.Message()))
	}
}

func (a *AssemblyTaskActor[T]) Process(ctx *actor.Context, msg *messages.AssemblyTaskMessage) {
	if a.task != msg.Task {
		panic(fmt.Sprintf("Actor \"%s\" received message for instance \"%s\"", a.task, msg.Task))
	}
	if step, ok := a.actor.Steps()[msg.Step]; ok {
		step(msg, ctx)
	} else {
		panic(fmt.Sprintf("Actor \"%s\" received unknown step \"%d\"", a.task, msg.Step))
	}
}
