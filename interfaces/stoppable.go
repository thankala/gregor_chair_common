package interfaces

import "github.com/anthdm/hollywood/actor"

type Stoppable interface {
	OnStopped(stoped actor.Stopped, ctx *actor.Context)
}
