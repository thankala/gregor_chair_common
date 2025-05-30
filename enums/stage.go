package enums

type Stage int

const (
	Initial           = Stage(NoneComponent)
	InitialBase       = Stage(NoneComponent | Base)
	LegsAttached      = Stage(NoneComponent | Base | Legs)
	CastorsAttached   = Stage(NoneComponent | Base | Legs | Castors)
	LiftAttached      = Stage(NoneComponent | Base | Legs | Castors | Lift)
	InitialSeat       = Stage(NoneComponent | Seat)
	SeatPlateAttached = Stage(NoneComponent | Seat | SeatPlate)
	SeatAttached      = Stage(NoneComponent | Base | Legs | Castors | Lift | Seat | SeatPlate)
	BackAttached      = Stage(NoneComponent | Base | Legs | Castors | Lift | Seat | SeatPlate | Back)
	LeftArmAttached   = Stage(NoneComponent | Base | Legs | Castors | Lift | Seat | SeatPlate | Back | LeftArm)
	RightArmAttached  = Stage(NoneComponent | Base | Legs | Castors | Lift | Seat | SeatPlate | Back | RightArm)
	Completed         = Stage(NoneComponent | Base | Legs | Castors | Lift | Seat | SeatPlate | Back | LeftArm | RightArm)
)

func (s Stage) String() string {
	switch s {
	case Initial:
		return "Initial"
	case InitialBase:
		return "InitialBase"
	case LegsAttached:
		return "LegsAttached"
	case SeatAttached:
		return "SeatAttached"
	case CastorsAttached:
		return "CastorsAttached"
	case LiftAttached:
		return "LiftAttached"
	case BackAttached:
		return "BackAttached"
	case LeftArmAttached:
		return "LeftArmAttached"
	case Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}
