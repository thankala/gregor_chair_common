package interfaces

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/messages"
)

type Coordinator[T any] interface {
	Coordinator() enums.Coordinator
	Process(ctx *actor.Context)
	RequestFixture(msg *messages.CoordinatorMessage)
	PlaceComponent(msg *messages.CoordinatorMessage)
	AttachComponent(msg *messages.CoordinatorMessage)
}
