package interfaces

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/events"
)

type Orchestrator[T any] interface {
	Orchestrator() enums.Task
	Process(ctx *actor.Context)
	FixtureRequested(event *events.OrchestratorEvent)
	ComponentPlaced(event *events.OrchestratorEvent)
	ComponentAttached(event *events.OrchestratorEvent)
}
