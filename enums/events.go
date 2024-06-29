package enums

type Event int

const (
	NoneEvent Event = iota
	AssemblyTaskEvent
	CoordinatorEvent
)
