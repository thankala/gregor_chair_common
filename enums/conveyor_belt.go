package enums

type ConveyorBelt int

const (
	ConveyorBelt1 ConveyorBelt = iota
	ConveyorBelt2
	ConveyorBelt3
)

func (c ConveyorBelt) String() string {
	return [...]string{"ConveyorBelt1", "ConveyorBelt2", "ConveyorBelt3"}[c]
}
