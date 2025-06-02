package states

import "github.com/thankala/gregor_chair_common/enums"

type RobotState struct {
	Item     enums.Component
	Task     enums.Task
	Position enums.Position
	Facing   string
}
