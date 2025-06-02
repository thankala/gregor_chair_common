package configuration

import "github.com/thankala/gregor_chair_common/enums"

type Location struct {
	X float64
	Y float64
	Z float64
	R float64
}

func NewLocation(x float64, y float64, z float64, r float64) Location {
	return Location{x, y, z, r}
}

type StorageConfiguration struct {
	Storage   enums.Storage
	Position  enums.Position
	Component enums.Component
	Location  Location
}

func NewStorageConfiguration(storage enums.Storage, position enums.Position, component enums.Component, location Location) *StorageConfiguration {
	return &StorageConfiguration{Storage: storage, Position: position, Component: component, Location: location}
}

type WorkbenchConfiguration struct {
	Workbench enums.Workbench
	Position  enums.Position
	Fixture   enums.Fixture
	Location  Location
}

func NewWorkbenchConfiguration(workbench enums.Workbench, position enums.Position, fixture enums.Fixture, location Location) *WorkbenchConfiguration {
	return &WorkbenchConfiguration{Workbench: workbench, Position: position, Fixture: fixture, Location: location}
}

type ConveyorBeltConfiguration struct {
	ConveyorBelt enums.ConveyorBelt
	Position     enums.Position
	Component    enums.Component
	Deposit      bool
	Location     Location
}

func NewConveyorBeltConfiguration(conveyorBelt enums.ConveyorBelt, position enums.Position, component enums.Component, deposit bool, location Location) *ConveyorBeltConfiguration {
	return &ConveyorBeltConfiguration{ConveyorBelt: conveyorBelt, Position: position, Component: component, Deposit: deposit, Location: location}
}

type RobotControllerConfiguration struct {
	Key           enums.Robot
	Storages      []StorageConfiguration      // Available storages
	Workbenches   []WorkbenchConfiguration    // Available workbenches
	ConveyorBelts []ConveyorBeltConfiguration // Available conveyor belts
}

type RobotConfigurationFunc func(configuration *RobotControllerConfiguration)

func WithRobotKey(key enums.Robot) RobotConfigurationFunc {
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
