package enums

type Fixture int

const (
	Fixture1 Fixture = iota
	Fixture2
	Fixture3
)

func (f Fixture) String() string {
	return [...]string{"Fixture1", "Fixture2", "Fixture3"}[f]
}

func (f Fixture) StringShort() string {
	return [...]string{"F1", "F2", "F3"}[f]
}
