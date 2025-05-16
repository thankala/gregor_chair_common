package events

import "github.com/thankala/gregor_chair_common/enums"

type AssemblyTaskEvent struct {
	Source      enums.Task      `json:"source"`
	Destination enums.Task      `json:"destination"`
	Step        enums.Step      `json:"step"`
	Component   enums.Component `json:"component"`
}
