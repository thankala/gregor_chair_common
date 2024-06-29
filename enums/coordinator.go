package enums

type Coordinator int

const (
	NoneCoordinator Coordinator = iota
	Coordinator1
	Coordinator2
)

func (c Coordinator) String() string {
	return [...]string{"NoneCoordinator", "Coordinator1", "Coordinator2"}[c]
}
