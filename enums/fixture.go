package enums

type Fixture string

const (
	NoneFixture Fixture = "NoneFixture"
	Fixture1    Fixture = "Fixture1"
	Fixture2    Fixture = "Fixture2"
	Fixture3    Fixture = "Fixture3"
)

func (f Fixture) String() string {
	return string(f)
}

func (f Fixture) StringShort() string {
	switch f {
	case Fixture1:
		return "F1"
	case Fixture2:
		return "F2"
	case Fixture3:
		return "F3"
	default:
		return "NoneFixture"
	}
}

type FixtureState string

const (
	Free       FixtureState = "FREE"
	Assembling FixtureState = "ASSEMBLING"
	Pending    FixtureState = "PENDING"
	Completed  FixtureState = "COMPLETED"
)

func (w FixtureState) String() string {
	return string(w)
}
