package interfaces

import "github.com/anthdm/hollywood/actor"

type Initializable interface {
	OnInitialized(initialized actor.Initialized, ctx *actor.Context)
}
