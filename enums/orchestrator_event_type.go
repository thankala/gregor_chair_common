package enums

type OrchestratorEventType string

const (
	None              OrchestratorEventType = "None"
	AssemblyStarted   OrchestratorEventType = "AssemblyStarted"
	FixtureRequested  OrchestratorEventType = "FixtureRequested"
	ComponentPlaced   OrchestratorEventType = "ComponentPlaced"
	ComponentAttached OrchestratorEventType = "ComponentAttached"
	ComponentPickedUp OrchestratorEventType = "ComponentPickedUp"
)

func (c OrchestratorEventType) String() string {
	return string(c)
}
