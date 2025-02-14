package models

import "github.com/thankala/gregor_chair_common/enums"

type Request struct {
	Task     enums.AssemblyTask
	Step     enums.Step
	Caller   string
	Expected []enums.Stage
	IsPickup bool
}
