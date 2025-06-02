package enums

type Robot string

const (
	Robot1 Robot = "Robot1"
	Robot2 Robot = "Robot2"
	Robot3 Robot = "Robot3"
)

func (r Robot) String() string {
	return string(r)
}

func (r Robot) StringShort() string {
	switch {
	case r == Robot1:
		return "R1"
	case r == Robot2:
		return "R2"
	case r == Robot3:
		return "R3"
	default:
		return ""
	}
}
