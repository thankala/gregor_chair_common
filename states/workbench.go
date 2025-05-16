package states

import (
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/models"
	"github.com/thankala/gregor_chair_common/utilities"
)

type FixtureState struct {
	Owner     enums.Task
	Component enums.Component
}

type WorkbenchState struct {
	Initialized bool
	Fixtures    map[enums.Fixture]FixtureState
	Requests    map[enums.Fixture]utilities.Queue[models.Request]
}
