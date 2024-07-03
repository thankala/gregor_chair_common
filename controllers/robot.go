package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/thankala/gregor_chair_common/configuration"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/interfaces"
	"github.com/thankala/gregor_chair_common/logger"
	"github.com/thankala/gregor_chair_common/states"
	"math/rand"
	"slices"
	"time"
)

type RobotController struct {
	storer        interfaces.Storer
	httpClient    interfaces.HttpClient
	configuration configuration.RobotControllerConfiguration
}

func NewRobotController(storer interfaces.Storer, httpClient interfaces.HttpClient, fns ...configuration.RobotConfigurationFunc) *RobotController {
	robotConfiguration := configuration.RobotControllerConfiguration{}
	for _, fn := range fns {
		fn(&robotConfiguration)
	}

	controller := &RobotController{
		storer:        storer,
		httpClient:    httpClient,
		configuration: robotConfiguration,
	}

	controller.resetState()
	controller.resetRobot()
	controller.releaseLock()
	return controller
}

func (c *RobotController) resetRobot() {
	if c.httpClient == nil {
		return
	}
	if _, err := c.httpClient.Post("/home", nil); err != nil {
		panic(err)
	}
}

func (c *RobotController) MoveToStorage(storage enums.Storage) {
	//time.Sleep(500 * time.Millisecond)
	storageConfiguration := c.getStorageConfiguration(storage)
	// TODO: Add machine integration
	state := c.loadState()
	state.Position = storageConfiguration.Position
	state.Facing = storageConfiguration.Storage.String()
	c.storeState(state)
	logger.Get().Info("Robot moved to storage", "Robot", c.configuration.Key, "Storage", storage.String())
}

func (c *RobotController) MoveToWorkbench(workbench enums.Workbench) {
	//time.Sleep(500 * time.Millisecond)
	workbenchConfiguration := c.getWorkbenchConfiguration(workbench)
	// TODO: Add machine integration
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/move/"+workbench.StringShort(), nil); err != nil {
			panic(err)
		}
	}
	state := c.loadState()
	state.Position = workbenchConfiguration.Position
	state.Facing = workbenchConfiguration.Workbench.String()
	c.storeState(state)
	logger.Get().Info("Robot moved to workbench", "Robot", c.configuration.Key, "Workbench", workbench.String())
}

func (c *RobotController) MoveToConveyorBelt(conveyorBelt enums.ConveyorBelt) {
	//time.Sleep(500 * time.Millisecond)
	conveyorBeltConfiguration := c.getConveyorBeltConfiguration(conveyorBelt)
	// TODO: Add machine integration
	state := c.loadState()
	state.Position = conveyorBeltConfiguration.Position
	state.Facing = conveyorBeltConfiguration.ConveyorBelt.String()
	c.storeState(state)
	logger.Get().Info("Robot moved to conveyor belt", "Robot", c.configuration.Key, "Conveyor Belt", conveyorBelt.String())
}

func (c *RobotController) PickAndPlace() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/pick_and_place", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("PickAndPlace executed", "Robot", c.configuration.Key)
	return
}

func (c *RobotController) PickAndInsert() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/pick_and_insert", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("PickAndInsert executed", "Robot", c.configuration.Key)
	return
}

func (c *RobotController) PickAndFlipAndPress() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/pick_and_flip_and_press", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("PickAndFlipAndPress executed", "Robot", c.configuration.Key)
	return
}

func (c *RobotController) ScrewPickAndFasten() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/screw_pick_and_fasten", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("ScrewPickAndFasten executed", "Robot", c.configuration.Key)
	return
}

// Pickup items

func (c *RobotController) PickupItemFromWorkbench(item enums.Component, workbench enums.Workbench) {
	//time.Sleep(500 * time.Millisecond)
	state := c.loadState()
	workbenchConfiguration := c.getWorkbenchConfiguration(workbench)
	if state.Facing != workbenchConfiguration.Workbench.String() || state.Position != workbenchConfiguration.Position {
		panic(fmt.Sprintf("Robot %s is facing %s and not specified coordinator_1 %s", c.configuration.Key, state.Facing, workbench.String()))
	}

	// TODO: Add machine integration
	state.Item = item
	c.storeState(state)
	logger.Get().Info("Robot picked up item", "Robot", c.configuration.Key, "Item", item, "From", workbench.String())
}

func (c *RobotController) PickupItemFromStorage(storage enums.Storage) {
	//time.Sleep(500 * time.Millisecond)
	state := c.loadState()
	storageConfiguration := c.getStorageConfiguration(storage)
	if state.Facing != storageConfiguration.Storage.String() || state.Position != storageConfiguration.Position {
		panic(fmt.Sprintf("Robot %s is facing %s and not specified storage %s", c.configuration.Key, state.Facing, storage.String()))
	}

	// TODO: Add machine integration
	state.Item = storageConfiguration.Component
	c.storeState(state)
	logger.Get().Info("Robot picked up item", "Robot", c.configuration.Key, "Item", storageConfiguration.Component, "From", storage.String())
}

func (c *RobotController) PickupItemFromConveyorBelt(conveyorBelt enums.ConveyorBelt) {
	//time.Sleep(500 * time.Millisecond)
	state := c.loadState()
	conveyorBeltConfiguration := c.getConveyorBeltConfiguration(conveyorBelt)
	if state.Facing != conveyorBeltConfiguration.ConveyorBelt.String() || state.Position != conveyorBeltConfiguration.Position {
		panic(fmt.Sprintf("Robot %s is facing %s and not specified belt %s", c.configuration.Key, state.Facing, conveyorBelt.String()))
	}

	// TODO: Add machine integration
	state.Item = conveyorBeltConfiguration.Component
	c.storeState(state)
	logger.Get().Info("Robot picked up item", "Robot", c.configuration.Key, "Item", conveyorBeltConfiguration.Component, "From", conveyorBelt.String())
}

// Pickup and deposit items
func (c *RobotController) ReleaseItem() enums.Component {
	//time.Sleep(500 * time.Millisecond)
	state := c.loadState()
	// TODO: Add machine integration
	item := state.Item
	state.Item = enums.NoneComponent
	c.storeState(state)
	logger.Get().Info("Robot released item", "Robot", c.configuration.Key, "Item", item)
	//fmt.Printf("Robot \"%s\" released item \"%v\"\n", c.configuration.Key, item)
	return item
}

// Task management

func (c *RobotController) WaitUntilFree() {
	for {
		duration := time.Duration(rand.Intn(100))
		time.Sleep(duration * time.Second)

		if c.GetCurrentTask() == enums.NoneAssemblyTask {
			break
		}
	}
}

func (c *RobotController) IsBusy() bool {
	return c.GetCurrentTask() != enums.NoneAssemblyTask
}

func (c *RobotController) ValidateCurrentTask(task enums.AssemblyTask) {
	if c.GetCurrentTask() != task {
		panic(fmt.Sprintf("Robot %s is not assigned to task %s", c.configuration.Key, task.String()))
	}
}

func (c *RobotController) GetCurrentTask() enums.AssemblyTask {
	state := c.loadState()
	return state.Task
}

func (c *RobotController) SetCurrentTask(task enums.AssemblyTask) error {
	res := c.acquireLock(task)
	if !res {
		return fmt.Errorf("robot %s is busy", c.configuration.Key)
	}

	state := c.loadState()
	if state.Task != enums.NoneAssemblyTask {
		//c.WaitUntilFree()
		return fmt.Errorf("robot %s is already assigned to task %s", c.configuration.Key, state.Task.String())
	}
	state.Task = task
	c.storeState(state)

	return nil
}

func (c *RobotController) ClearCurrentTask() {
	state := c.loadState()
	state.Task = enums.NoneAssemblyTask
	c.storeState(state)
	c.releaseLock()
}

// State Management
func (c *RobotController) acquireLock(task enums.AssemblyTask) bool {
	res, err := c.storer.AcquireLock(fmt.Sprintf("%s-lock", c.configuration.Key), task.String())
	if err != nil {
		panic("Failed to acquire lock")
	}
	return res
}

func (c *RobotController) releaseLock() {
	err := c.storer.ReleaseLock(fmt.Sprintf("%s-lock", c.configuration.Key))
	if err != nil {
		panic("Failed to release lock")
	}
}

func (c *RobotController) resetState() {
	c.storeState(states.RobotState{
		Position: enums.NonePosition,
		Facing:   enums.NoneWorkbench.String(),
		Item:     enums.NoneComponent,
		Task:     enums.NoneAssemblyTask,
	})
}

func (c *RobotController) loadState() states.RobotState {
	var state states.RobotState
	v, err := c.storer.Load(c.configuration.Key)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(v, &state); err != nil {
		panic(err)
	}
	return state
}

func (c *RobotController) storeState(state states.RobotState) {
	v, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	if err := c.storer.Store(c.configuration.Key, v); err != nil {
		panic(err)
	}
}

// Configuration Management

func (c *RobotController) Key() string {
	return c.configuration.Key
}

func (c *RobotController) getStorageConfiguration(storage enums.Storage) configuration.StorageConfiguration {
	storageIndex := slices.IndexFunc(c.configuration.Storages, func(storageConfiguration configuration.StorageConfiguration) bool {
		return storageConfiguration.Storage == storage
	})
	if storageIndex == -1 {
		panic("Invalid storage")
	}
	return c.configuration.Storages[storageIndex]
}

func (c *RobotController) getWorkbenchConfiguration(workbench enums.Workbench) configuration.WorkbenchConfiguration {
	workbenchIndex := slices.IndexFunc(c.configuration.Workbenches, func(workbenchConfiguration configuration.WorkbenchConfiguration) bool {
		return workbenchConfiguration.Workbench == workbench
	})
	if workbenchIndex == -1 {
		panic("Invalid coordinator_1")
	}
	return c.configuration.Workbenches[workbenchIndex]
}

func (c *RobotController) getConveyorBeltConfiguration(conveyorBelt enums.ConveyorBelt) configuration.ConveyorBeltConfiguration {
	conveyorBeltIndex := slices.IndexFunc(c.configuration.ConveyorBelts, func(conveyorBeltConfiguration configuration.ConveyorBeltConfiguration) bool {
		return conveyorBeltConfiguration.ConveyorBelt == conveyorBelt
	})
	if conveyorBeltIndex == -1 {
		panic("Invalid storage")
	}
	return c.configuration.ConveyorBelts[conveyorBeltIndex]
}
