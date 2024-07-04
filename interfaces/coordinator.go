package interfaces

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/messages"
)

type Coordinator[T any] interface {
	Coordinator() enums.Coordinator
	Process(ctx *actor.Context)
	FixtureRequested(msg *messages.CoordinatorMessage)
	ComponentPlaced(msg *messages.CoordinatorMessage)
	ComponentAttached(msg *messages.CoordinatorMessage)
}
