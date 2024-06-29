package enums

type Workbench int

const (
	NoneWorkbench Workbench = iota
	Workbench1
	Workbench2
)

func (w Workbench) String() string {
	return [...]string{"NoneComponent", "Workbench1", "Workbench2"}[w]
}

func (w Workbench) StringShort() string {
	return [...]string{"NoneComponent", "W1", "W2"}[w]
}
