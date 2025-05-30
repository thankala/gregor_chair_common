package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/thankala/gregor_chair_common/configuration"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/interfaces"
	"github.com/thankala/gregor_chair_common/logger"
	"github.com/thankala/gregor_chair_common/states"
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
	storageConfiguration := c.getStorageConfiguration(storage)
	state := c.loadState()
	state.Position = storageConfiguration.Position
	state.Facing = storageConfiguration.Storage.String()
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/move/"+storage.StringShort(), nil); err != nil {
			panic(err)
		}
	}
	c.storeState(state)
	logger.Get().Info("Robot moved to storage", "Robot", c.configuration.Key, "Storage", storage.String(), "Task", c.GetCurrentTask())
}

func (c *RobotController) MoveToWorkbench(workbench enums.Workbench) {
	workbenchConfiguration := c.getWorkbenchConfiguration(workbench)
	state := c.loadState()
	state.Position = workbenchConfiguration.Position
	state.Facing = workbenchConfiguration.Workbench.String()
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/move/"+workbench.StringShort(), nil); err != nil {
			panic(err)
		}
	}
	c.storeState(state)
	logger.Get().Info("Robot moved to workbench", "Robot", c.configuration.Key, "Workbench", workbench.String(), "Task", c.GetCurrentTask())
}

func (c *RobotController) MoveToConveyorBelt(conveyorBelt enums.ConveyorBelt) {
	conveyorBeltConfiguration := c.getConveyorBeltConfiguration(conveyorBelt)
	state := c.loadState()
	state.Position = conveyorBeltConfiguration.Position
	state.Facing = conveyorBeltConfiguration.ConveyorBelt.String()
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/move/"+conveyorBelt.StringShort(), nil); err != nil {
			panic(err)
		}
	}
	c.storeState(state)
	logger.Get().Info("Robot moved to conveyor belt", "Robot", c.configuration.Key, "Conveyor Belt", conveyorBelt.String(), "Task", c.GetCurrentTask())
}

func (c *RobotController) Pick() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/grip", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("Grip executed", "Robot", c.configuration.Key, "Task", c.GetCurrentTask())
}

func (c *RobotController) Place() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/ungrip", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("Ungrip executed", "Robot", c.configuration.Key, "Task", c.GetCurrentTask())
}

func (c *RobotController) Screw() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/screw", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("Screw executed", "Robot", c.configuration.Key, "Task", c.GetCurrentTask())
}

func (c *RobotController) Flip() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/flip", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("Flip executed", "Robot", c.configuration.Key, "Task", c.GetCurrentTask())
}

func (c *RobotController) Press() {
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/composite/press", nil); err != nil {
			panic(err)
		}
	}
	logger.Get().Info("Press executed", "Robot", c.configuration.Key, "Task", c.GetCurrentTask())
}

// Pickup items

func (c *RobotController) PickupItemFromWorkbench(item enums.Component, workbench enums.Workbench) {
	state := c.loadState()
	workbenchConfiguration := c.getWorkbenchConfiguration(workbench)
	if state.Facing != workbenchConfiguration.Workbench.String() || state.Position != workbenchConfiguration.Position {
		panic(fmt.Sprintf("Robot %s is facing %s and not specified coordinator_1 %s", c.configuration.Key, state.Facing, workbench.String()))
	}

	body := map[string]bool{
		"enable_grip":    true,
		"enable_control": true,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/grip", bytes.NewBuffer(jsonData)); err != nil {
			panic(err)
		}
	}

	state.Item = item
	c.storeState(state)
	logger.Get().Info("Robot picked up item", "Robot", c.configuration.Key, "Item", item.String(), "From", workbench.String(), "Task", c.GetCurrentTask())
}

func (c *RobotController) PickupItemFromStorage(storage enums.Storage) {
	state := c.loadState()
	storageConfiguration := c.getStorageConfiguration(storage)
	if state.Facing != storageConfiguration.Storage.String() || state.Position != storageConfiguration.Position {
		panic(fmt.Sprintf("Robot %s is facing %s and not specified storage %s", c.configuration.Key, state.Facing, storage.String()))
	}

	body := map[string]bool{
		"enable_grip":    true,
		"enable_control": true,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/grip", bytes.NewBuffer(jsonData)); err != nil {
			panic(err)
		}
	}

	state.Item = storageConfiguration.Component
	c.storeState(state)
	logger.Get().Info("Robot picked up item", "Robot", c.configuration.Key, "Item", storageConfiguration.Component.String(), "From", storage.String(), "Task", c.GetCurrentTask())
}

func (c *RobotController) PickupItemFromConveyorBelt(conveyorBelt enums.ConveyorBelt) {
	state := c.loadState()
	conveyorBeltConfiguration := c.getConveyorBeltConfiguration(conveyorBelt)
	if state.Facing != conveyorBeltConfiguration.ConveyorBelt.String() || state.Position != conveyorBeltConfiguration.Position {
		panic(fmt.Sprintf("Robot %s is facing %s and not specified belt %s", c.configuration.Key, state.Facing, conveyorBelt.String()))
	}

	body := map[string]bool{
		"enable_grip":    true,
		"enable_control": true,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/primitive/grip", bytes.NewBuffer(jsonData)); err != nil {
			panic(err)
		}
	}

	state.Item = conveyorBeltConfiguration.Component
	c.storeState(state)
	logger.Get().Info("Robot picked up item", "Robot", c.configuration.Key, "Item", conveyorBeltConfiguration.Component.String(), "From", conveyorBelt.String(), "Task", c.GetCurrentTask())
}

// Pickup and deposit items

func (c *RobotController) ReleaseItem() enums.Component {
	state := c.loadState()
	item := state.Item
	logger.Get().Info("Robot released item", "Robot", c.configuration.Key, "Item", item.String(), "Task", c.GetCurrentTask())
	// if c.httpClient != nil {
	// 	if _, err := c.httpClient.Post("/composite/place", nil); err != nil {
	// 		panic(err)
	// 	}
	// }
	state.Item = enums.NoneComponent
	c.storeState(state)
	return item
}

// Task management

func (c *RobotController) WaitUntilFree() {
	for {
		duration := time.Duration(rand.Intn(100))
		time.Sleep(duration * time.Second)

		if c.GetCurrentTask() == enums.NoneTask {
			break
		}
	}
}

func (c *RobotController) IsBusy() bool {
	return c.GetCurrentTask() != enums.NoneTask
}

func (c *RobotController) ValidateCurrentTask(task enums.Task) {
	if c.GetCurrentTask() != task {
		panic(fmt.Sprintf("Robot %s is not assigned to task %s", c.configuration.Key, task))
	}
}

func (c *RobotController) GetCurrentTask() enums.Task {
	state := c.loadState()
	return state.Task
}

func (c *RobotController) SetCurrentTask(task enums.Task) error {
	res, err := c.acquireLock(task)

	if err != nil {
		return err
	}
	if !res {
		return fmt.Errorf("robot %s is busy", c.configuration.Key)
	}

	state := c.loadState()
	if state.Task != enums.NoneTask {
		//c.WaitUntilFree()
		return fmt.Errorf("robot %s is already assigned to task %s", c.configuration.Key, state.Task)
	}
	state.Task = task
	c.storeState(state)
	logger.Get().Info("Robot started task", "Robot", c.configuration.Key, "Task", c.GetCurrentTask())
	return nil
}

func (c *RobotController) ClearCurrentTask() {
	state := c.loadState()
	logger.Get().Info("Robot ended task", "Robot", c.configuration.Key, "Task", state.Task.String())
	state.Task = enums.NoneTask
	c.storeState(state)
	c.releaseLock()
}

// State Management
func (c *RobotController) acquireLock(task enums.Task) (bool, error) {
	res, err := c.storer.AcquireLock(fmt.Sprintf("%s-lock", c.configuration.Key), string(task))
	if err != nil {
		return false, fmt.Errorf("failed to acquire lock: %s", err)
	}
	return res, nil
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
		Task:     enums.NoneTask,
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
		panic("Invalid workbench")
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
