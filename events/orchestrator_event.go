package events

import "github.com/thankala/gregor_chair_common/enums"

type OrchestratorEvent struct {
	Source      enums.Task                  `json:"source"`
	Destination enums.Task                  `json:"destination"`
	Type        enums.OrchestratorEventType `json:"type"`
	Step        enums.Step                  `json:"step"`
	Caller      enums.Robot                 `json:"caller"`
	Workbench   enums.Workbench             `json:"workbench"`
	Fixture     enums.Fixture               `json:"fixture"`
	Expected    []enums.Stage               `json:"expected"`
	Component   enums.Component             `json:"component"`
}
