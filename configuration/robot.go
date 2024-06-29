package configuration

import "github.com/thankala/gregor_chair_common/enums"

type StorageConfiguration struct {
	Storage   enums.Storage
	Position  enums.Position
	Component enums.Component
}

func NewStorageConfiguration(storage enums.Storage, position enums.Position, component enums.Component) *StorageConfiguration {
	return &StorageConfiguration{Storage: storage, Position: position, Component: component}
}

type WorkbenchConfiguration struct {
	Workbench enums.Workbench
	Position  enums.Position
	Fixture   enums.Fixture
}

func NewWorkbenchConfiguration(workbench enums.Workbench, position enums.Position, fixture enums.Fixture) *WorkbenchConfiguration {
	return &WorkbenchConfiguration{Workbench: workbench, Position: position, Fixture: fixture}
}

type ConveyorBeltConfiguration struct {
	ConveyorBelt enums.ConveyorBelt
	Position     enums.Position
	Component    enums.Component
	Deposit      bool
}

func NewConveyorBeltConfiguration(conveyorBelt enums.ConveyorBelt, position enums.Position, component enums.Component, deposit bool) *ConveyorBeltConfiguration {
	return &ConveyorBeltConfiguration{ConveyorBelt: conveyorBelt, Position: position, Component: component, Deposit: deposit}
}

type RobotControllerConfiguration struct {
	Key           string
	Storages      []StorageConfiguration      // Available storages
	Workbenches   []WorkbenchConfiguration    // Available workbenches
	ConveyorBelts []ConveyorBeltConfiguration // Available conveyor belts
}

type RobotConfigurationFunc func(configuration *RobotControllerConfiguration)

func WithRobotKey(key string) RobotConfigurationFunc {
	return func(configuration *RobotControllerConfiguration) {
		configuration.Key = key
	}
}

func WithStorages(storages ...StorageConfiguration) RobotConfigurationFunc {
	return func(configuration *RobotControllerConfiguration) {
		configuration.Storages = storages
	}
}

func WithWorkbenches(workbenches ...WorkbenchConfiguration) RobotConfigurationFunc {
	return func(configuration *RobotControllerConfiguration) {
		configuration.Workbenches = workbenches
	}
}

func WithConveyorBelts(conveyorBelts ...ConveyorBeltConfiguration) RobotConfigurationFunc {
	return func(configuration *RobotControllerConfiguration) {
		configuration.ConveyorBelts = conveyorBelts
	}
}
