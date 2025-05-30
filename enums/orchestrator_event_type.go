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
