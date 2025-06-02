package enums

type Component int

const (
	NoneComponent    Component = 0
	Base             Component = 1
	Legs             Component = 2
	Castors          Component = 4
	Lift             Component = 8
	Seat             Component = 16
	SeatPlate        Component = 32
	SeatAndSeatPlate Component = 48
	Back             Component = 64
	LeftArm          Component = 128
	RightArm         Component = 256
)

func (c *Component) Stage() Stage {
	return Stage(*c)
}

func (c *Component) String() string {
	switch *c {
	case Legs:
		return "Legs"
	case Base:
		return "Base"
	case Castors:
		return "Castors"
	case Lift:
		return "Lift"
	case Seat:
		return "Seat"
	case SeatPlate:
		return "SeatPlate"
	case SeatAndSeatPlate:
		return "SeatAndSeatPlate"
	case Back:
		return "Back"
	case LeftArm:
		return "LeftArm"
	case RightArm:
		return "RightArm"
	default:
		return "None"
	}
}
