package interfaces

import "github.com/anthdm/hollywood/actor"

type Startable interface {
	OnStarted(started actor.Started, ctx *actor.Context)
}
