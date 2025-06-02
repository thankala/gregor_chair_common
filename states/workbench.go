package states

import (
	"github.com/thankala/gregor_chair_common/enums"
)

type FixtureState struct {
	Owner     enums.Task
	Component enums.Component
}

type WorkbenchState struct {
	Initialized bool
	Fixtures    map[enums.Fixture]FixtureState
}
