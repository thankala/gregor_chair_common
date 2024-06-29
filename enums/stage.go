package enums

type Stage int

const (
	Initial            = Stage(NoneComponent)
	LegsAttached       = Stage(NoneComponent | Legs)
	BaseAttached       = Stage(NoneComponent | Legs | Base)
	CastorsAttached    = Stage(NoneComponent | Legs | Base | Castors)
	LiftAttached       = Stage(NoneComponent | Legs | Base | Castors | Lift)
	InitialSeat        = Stage(NoneComponent | Seat)
	SeatPlateAttached  = Stage(NoneComponent | Seat | SeatPlate)
	SeatScrewsAttached = Stage(NoneComponent | Seat | SeatPlate | SeatScrews)
	SeatAttached       = Stage(NoneComponent | Legs | Base | Castors | Lift | Seat | SeatPlate | SeatScrews)
	ScrewsAttached     = Stage(NoneComponent | Legs | Base | Castors | Lift | Seat | SeatPlate | SeatScrews | Screws)
	BackAttached       = Stage(NoneComponent | Legs | Base | Castors | Lift | Seat | SeatPlate | SeatScrews | Screws | Back)
	LeftArmAttached    = Stage(NoneComponent | Legs | Base | Castors | Lift | Seat | SeatPlate | SeatScrews | Screws | Back | LeftArm)
	RightArmAttached   = Stage(NoneComponent | Legs | Base | Castors | Lift | Seat | SeatPlate | SeatScrews | Screws | Back | RightArm)
	Completed          = Stage(NoneComponent | Legs | Base | Castors | Lift | Seat | SeatPlate | SeatScrews | Screws | Back | LeftArm | RightArm)
)

func (s Stage) String() string {
	switch s {
	case Initial:
		return "Initial"
	case LegsAttached:
		return "LegsAttached"
	case BaseAttached:
		return "BaseAttached"
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
