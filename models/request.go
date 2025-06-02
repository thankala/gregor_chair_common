package models

import "github.com/thankala/gregor_chair_common/enums"

type Request struct {
	Task     enums.Task
	Step     enums.Step
	Type     enums.OrchestratorEventType
	Caller   string
	Expected []enums.Stage
	IsPickup bool
}
