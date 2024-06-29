package messages

import "github.com/thankala/gregor_chair_common/enums"

type CoordinatorMessage struct {
	Event       enums.Event                  `json:"event"`
	Source      string                       `json:"source"`
	Destination string                       `json:"destination"`
	Task        enums.AssemblyTask           `json:"task"`
	Type        enums.CoordinatorMessageType `json:"type"`
	Item        enums.Component              `json:"item"`
	Step        enums.Step                   `json:"step"`
	Caller      string                       `json:"caller"`
	Fixture     enums.Fixture                `json:"fixture"`
	Expected    []enums.Stage                `json:"expected"`
	IsPickup    bool                         `json:"is_pickup"`
	Component   enums.Component              `json:"component"`
}
