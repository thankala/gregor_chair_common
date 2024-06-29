package messages

import "github.com/thankala/gregor_chair_common/enums"

type AssemblyTaskMessage struct {
	Event       enums.Event        `json:"event"`
	Source      string             `json:"source"`
	Destination string             `json:"destination"`
	Task        enums.AssemblyTask `json:"task"`
	Step        enums.Step         `json:"step"`
	Component   enums.Component    `json:"component"`
}
