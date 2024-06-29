package configuration

import "github.com/thankala/gregor_chair_common/enums"

type FixtureConfiguration struct {
	Fixture     enums.Fixture
	Subscribers []string
}

func NewFixtureConfiguration(fixture enums.Fixture, subscribers []string) *FixtureConfiguration {
	return &FixtureConfiguration{Fixture: fixture, Subscribers: subscribers}
}

type WorkbenchControllerConfiguration struct {
	Key          string
	Fixtures     []FixtureConfiguration
	StateMapping map[enums.Fixture]map[enums.Stage]string
}

type WorkbenchControllerConfigurationFunc func(configuration *WorkbenchControllerConfiguration)

func WithWorkbenchKey(key string) WorkbenchControllerConfigurationFunc {
	return func(configuration *WorkbenchControllerConfiguration) {
		configuration.Key = key
	}
}

func WithFixture(fixtures ...FixtureConfiguration) WorkbenchControllerConfigurationFunc {
	return func(configuration *WorkbenchControllerConfiguration) {
		configuration.Fixtures = fixtures
	}
}

func WithStateMapping(mapping map[enums.Fixture]map[enums.Stage]string) WorkbenchControllerConfigurationFunc {
	return func(configuration *WorkbenchControllerConfiguration) {
		configuration.StateMapping = mapping
	}
}
