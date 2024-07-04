package enums

type CoordinatorMessageType int

const (
	None CoordinatorMessageType = iota
	ComponentPlaced
	FixtureRequested
	ComponentAttached
)
