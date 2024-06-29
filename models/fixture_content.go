package models

import "github.com/thankala/gregor_chair_common/enums"

type FixtureContent struct {
	Owner     enums.AssemblyTask
	Fixture   enums.Fixture
	Component enums.Component
}
