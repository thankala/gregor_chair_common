package interfaces

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/events"
)

type StepHandler[T any] func(event *events.AssemblyTaskEvent, ctx *actor.Context)
type StepHandlers[T any] map[enums.Step]StepHandler[T]

type AssemblyTask[T any] interface {
	Task() enums.Task
	Steps() StepHandlers[T]
}
