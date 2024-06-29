package enums

type CoordinatorMessageType int

const (
	None CoordinatorMessageType = iota
	PlaceComponent
	RequestPickup
	RequestFixture
	AttachComponent
)
