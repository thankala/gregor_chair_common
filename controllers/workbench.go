package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/thankala/gregor_chair_common/configuration"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/interfaces"
	"github.com/thankala/gregor_chair_common/models"
	"github.com/thankala/gregor_chair_common/states"
	"github.com/thankala/gregor_chair_common/utilities"
	"slices"
)

type WorkbenchController struct {
	storer        interfaces.Storer
	httpClient    interfaces.HttpClient
	configuration configuration.WorkbenchControllerConfiguration
}

func NewWorkbenchController(storer interfaces.Storer, httpClient interfaces.HttpClient, fns ...configuration.WorkbenchControllerConfigurationFunc) *WorkbenchController {
	workbenchControllerConfiguration := configuration.WorkbenchControllerConfiguration{}
	for _, fn := range fns {
		fn(&workbenchControllerConfiguration)
	}
	controller := &WorkbenchController{
		storer:        storer,
		configuration: workbenchControllerConfiguration,
		httpClient:    httpClient,
	}
	controller.resetState()
	return controller
}

// Fixture Management

func (c *WorkbenchController) String() string {
	return c.configuration.Key
}

func (c *WorkbenchController) getFixturesContentInternal(state states.WorkbenchState) []models.FixtureContent {
	fixtures := make([]models.FixtureContent, len(c.configuration.Fixtures))
	for i, fixtureConfiguration := range c.configuration.Fixtures {
		fixtureState := state.Fixtures[fixtureConfiguration.Fixture]
		fixtures[i] = models.FixtureContent{
			Owner:     fixtureState.Owner,
			Fixture:   fixtureConfiguration.Fixture,
			Component: fixtureState.Component,
		}
	}
	return fixtures
}

func (c *WorkbenchController) canRotateInternal(initialized bool, fixtures []models.FixtureContent) bool {
	// If the coordinator_1 3 fixtures, it can rotate
	has3Fixtures := len(fixtures) == 3

	// If the first fixture is at stage BaseAttached and the other fixtures are at stage Initial, it can rotate
	isInitialConfiguration := !initialized && fixtures[0].Component.Stage() == enums.BaseAttached && fixtures[1].Component.Stage() == enums.Initial && fixtures[2].Component.Stage() == enums.Initial

	// Or if the first fixture is at stage BaseAttached ant the second fixture is at stage ScrewsAttached and the third fixture is at stage Initial, it can rotate
	isEveryConfiguration := initialized && fixtures[0].Component.Stage() == enums.BaseAttached && fixtures[1].Component.Stage() == enums.ScrewsAttached && fixtures[2].Component.Stage() == enums.Initial

	return has3Fixtures && (isInitialConfiguration || isEveryConfiguration)
}

func (c *WorkbenchController) GetFixturesContent() []models.FixtureContent {
	state := c.loadState()
	return c.getFixturesContentInternal(state)
}

func (c *WorkbenchController) CanRotate() bool {
	state := c.loadState()
	fixtures := c.getFixturesContentInternal(state)
	return c.canRotateInternal(state.Initialized, fixtures)
}

func (c *WorkbenchController) SetFixtureOwner(task enums.AssemblyTask, caller string, fixture enums.Fixture) {
	state := c.loadState()
	fixtureConfiguration := c.getFixtureConfiguration(caller, fixture)
	fixtureState := state.Fixtures[fixtureConfiguration.Fixture]
	fixtureState.Owner = task
	state.Fixtures[fixtureConfiguration.Fixture] = fixtureState
	c.storeState(state)
}

func (c *WorkbenchController) RotateFixtures() []models.FixtureContent {
	state := c.loadState()
	fixtures := c.getFixturesContentInternal(state)
	if !c.canRotateInternal(state.Initialized, fixtures) {
		return nil
	}

	// TODO: Add machine integration
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/rotate", nil); err != nil {
			panic(err)
		}
	}

	fixture1State := states.FixtureState{Owner: enums.NoneAssemblyTask, Component: enums.NoneComponent}
	fixture2State := state.Fixtures[fixtures[0].Fixture]
	fixture3State := state.Fixtures[fixtures[1].Fixture]

	state.Fixtures[fixtures[0].Fixture] = fixture1State
	state.Fixtures[fixtures[1].Fixture] = fixture2State
	state.Fixtures[fixtures[2].Fixture] = fixture3State

	if !state.Initialized {
		state.Initialized = true
	}

	c.storeState(state)
	return c.getFixturesContentInternal(state)
}

// Item Management

func (c *WorkbenchController) GetItem(task enums.AssemblyTask, caller string, fixture enums.Fixture) enums.Component {
	state := c.loadState()
	fixtureConfiguration := c.getFixtureConfiguration(caller, fixture)
	fixtureState := state.Fixtures[fixtureConfiguration.Fixture]

	if fixtureState.Owner != task {
		panic(fmt.Sprintf("Fixture \"%s\" is not owned by task \"%s\"", fixtureConfiguration.Fixture.String(), task))
	}

	// TODO: Add machine integration
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/fixtures/"+fixture.StringShort(), map[string]interface{}{
			"state": "FREE",
		}); err != nil {
			panic(err)
		}
	}

	component := fixtureState.Component
	fixtureState.Component = enums.NoneComponent
	state.Fixtures[fixtureConfiguration.Fixture] = fixtureState
	c.storeState(state)
	return component
}

func (c *WorkbenchController) SetItem(task enums.AssemblyTask, caller string, fixture enums.Fixture, component enums.Component) {
	state := c.loadState()
	fixtureConfiguration := c.getFixtureConfiguration(caller, fixture)
	fixtureState := state.Fixtures[fixtureConfiguration.Fixture]

	if fixtureState.Owner != task {
		panic(fmt.Sprintf("Fixture \"%s\" is not owned by task \"%s\"", fixtureConfiguration.Fixture.String(), task))
	}

	// TODO: Add machine integration
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/fixtures/"+fixture.StringShort(), map[string]interface{}{
			"state": "PENDING",
		}); err != nil {
			panic(err)
		}
	}

	fixtureState.Component = component
	state.Fixtures[fixtureConfiguration.Fixture] = fixtureState
	c.storeState(state)
}

func (c *WorkbenchController) AttachItem(task enums.AssemblyTask, caller string, fixture enums.Fixture, component enums.Component) {
	state := c.loadState()
	fixtureConfiguration := c.getFixtureConfiguration(caller, fixture)
	fixtureState := state.Fixtures[fixtureConfiguration.Fixture]

	if fixtureState.Owner != task {
		panic(fmt.Sprintf("Fixture \"%s\" is not owned by task \"%s\". It is owned by \"%s\"", fixtureConfiguration.Fixture.String(), task, fixtureState.Owner.String()))
	}

	// TODO: Add machine integration
	if c.httpClient != nil {
		if _, err := c.httpClient.Post("/fixtures/"+fixture.StringShort(), map[string]interface{}{
			"state": "ASSEMBLING",
		}); err != nil {
			panic(err)
		}
	}

	fixtureState.Component |= component
	state.Fixtures[fixtureConfiguration.Fixture] = fixtureState
	c.storeState(state)
}

func (c *WorkbenchController) RemoveCompletedItem() {
	state := c.loadState()
	fixtureState := state.Fixtures[enums.Fixture3]
	if fixtureState.Component.Stage() == enums.Completed {
		fixtureState.Component = enums.NoneComponent
		state.Fixtures[enums.Fixture3] = fixtureState
		c.storeState(state)
	}
}

// Request Management
func (c *WorkbenchController) PushRequest(request models.Request, fixture enums.Fixture) {
	state := c.loadState()
	fixtureConfiguration := c.getFixtureConfiguration(request.Caller, fixture)

	queue := state.Requests[fixtureConfiguration.Fixture]
	queue.Push(request)
	state.Requests[fixtureConfiguration.Fixture] = queue
	c.storeState(state)
}

func (c *WorkbenchController) PopRequest(fixture enums.Fixture) *models.Request {
	state := c.loadState()

	queue := state.Requests[fixture]
	request := queue.Pop()
	state.Requests[fixture] = queue

	c.storeState(state)
	return request
}

func (c *WorkbenchController) PeekRequest(fixture enums.Fixture) *models.Request {
	state := c.loadState()

	queue := state.Requests[fixture]
	return queue.Peek()
}

func (c *WorkbenchController) PeekAllRequests(fixture enums.Fixture) []models.Request {
	state := c.loadState()

	queue := state.Requests[fixture]
	return *queue.PeekAll()
}

func (c *WorkbenchController) SetLEDs(fixtures []models.FixtureContent) {
	if c.httpClient == nil {
		return
	}
	for _, fixture := range fixtures {
		state := c.configuration.StateMapping[fixture.Fixture][fixture.Component.Stage()]
		if _, err := c.httpClient.Post("/fixtures/"+fixture.Fixture.StringShort(), map[string]interface{}{
			"state": state,
		}); err != nil {
			panic(err)
		}
	}
}

// State Management
func (c *WorkbenchController) resetState() {
	state := states.WorkbenchState{}
	state.Fixtures = make(map[enums.Fixture]states.FixtureState)
	state.Requests = make(map[enums.Fixture]utilities.Queue[models.Request])
	for _, fixtureConfiguration := range c.configuration.Fixtures {
		state.Fixtures[fixtureConfiguration.Fixture] = states.FixtureState{
			Owner:     enums.NoneAssemblyTask,
			Component: enums.NoneComponent,
		}
		state.Requests[fixtureConfiguration.Fixture] = []models.Request{}
	}
	c.storeState(state)
}

func (c *WorkbenchController) loadState() states.WorkbenchState {
	var state states.WorkbenchState
	v, err := c.storer.Load(c.configuration.Key)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(v, &state); err != nil {
		panic(err)
	}
	return state
}

func (c *WorkbenchController) storeState(state states.WorkbenchState) {
	v, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	if err := c.storer.Store(c.configuration.Key, v); err != nil {
		panic(err)
	}
}

// Configuration Management
func (c *WorkbenchController) getFixtureConfiguration(caller string, fixture enums.Fixture) configuration.FixtureConfiguration {
	fixtureIndex := slices.IndexFunc(c.configuration.Fixtures, func(fixtureConfiguration configuration.FixtureConfiguration) bool {
		return fixtureConfiguration.Fixture == fixture && slices.Contains(fixtureConfiguration.Subscribers, caller)
	})
	if fixtureIndex == -1 {
		panic("Invalid fixture")
	}

	return c.configuration.Fixtures[fixtureIndex]
}
