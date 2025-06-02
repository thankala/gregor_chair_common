package enums

type ConveyorBelt string

const (
	ConveyorBelt1 ConveyorBelt = "ConveyorBelt1"
	ConveyorBelt2 ConveyorBelt = "ConveyorBelt2"
	ConveyorBelt3 ConveyorBelt = "ConveyorBelt3"
)

func (cb ConveyorBelt) String() string {
	return string(cb)
}

func (cb ConveyorBelt) StringShort() string {
	switch {
	case cb == ConveyorBelt1:
		return "CB1"
	case cb == ConveyorBelt2:
		return "CB2"
	case cb == ConveyorBelt3:
		return "CB3"
	default:
		return "N"
	}
}
