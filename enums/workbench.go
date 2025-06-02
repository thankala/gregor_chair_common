package enums

type Workbench string

const (
	NoneWorkbench Workbench = "NoneWorkbench"
	Workbench1    Workbench = "Workbench1"
	Workbench2    Workbench = "Workbench2"
)

func (w Workbench) String() string {
	return string(w)
}

func (w Workbench) StringShort() string {
	switch {
	case w == Workbench1:
		return "W1"
	case w == Workbench2:
		return "W2"
	default:
		return "N"
	}
}
