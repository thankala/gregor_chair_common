package interfaces

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/events"
)

type Orchestrator[T any] interface {
	Orchestrator() enums.Task
	Process(ctx *actor.Context, event *events.OrchestratorEvent)
	StartAssembly(ctx *actor.Context, event *events.OrchestratorEvent)
}
