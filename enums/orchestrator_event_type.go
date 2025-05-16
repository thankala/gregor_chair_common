package enums

type OrchestratorEventType string

const (
	None              OrchestratorEventType = "None"
	ComponentPlaced   OrchestratorEventType = "ComponentPlaced"
	FixtureRequested  OrchestratorEventType = "FixtureRequested"
	ComponentAttached OrchestratorEventType = "ComponentAttached"
)

func (c OrchestratorEventType) String() string {
	return string(c)
}

// type CoordinatorMessageType int

// const (
// 	None CoordinatorMessageType = iota
// 	ComponentPlaced
// 	FixtureRequested
// 	ComponentAttached
// )

// func (c CoordinatorMessageType) String() string {
// 	return [...]string{"None", "ComponentPlaced", "FixtureRequested", "ComponentAttached"}[c]
// }
