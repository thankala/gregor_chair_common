package configuration

import (
	"github.com/thankala/gregor_chair_common/enums"
)

type FixtureConfiguration struct {
	Fixture     enums.Fixture
	Subscribers []string
}

type FixtureConfigurationFunc func(configuration *FixtureConfiguration)

func NewFixtureConfiguration(fixture enums.Fixture, subscribers []string) *FixtureConfiguration {
	return &FixtureConfiguration{Fixture: fixture, Subscribers: subscribers}
}

type WorkbenchControllerConfiguration struct {
	Key      enums.Workbench
	Fixtures []FixtureConfiguration
}

type WorkbenchControllerConfigurationFunc func(configuration *WorkbenchControllerConfiguration)

func WithWorkbenchKey(key enums.Workbench) WorkbenchControllerConfigurationFunc {
	return func(configuration *WorkbenchControllerConfiguration) {
		configuration.Key = key
	}
}

func WithFixture(fixtures ...FixtureConfiguration) WorkbenchControllerConfigurationFunc {
	return func(configuration *WorkbenchControllerConfiguration) {
		configuration.Fixtures = fixtures
	}
}
