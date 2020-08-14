package models

import (
	"time"
)

// FiniteStateMachine used to implement fsm
type FiniteStateMachine struct {
	state            string
	handlers         map[string]*StateHandler
	activeAgent      Base
	currentTurnEndAt time.Time
}

// MakeFiniteStateMachine used to instantiate fsm
func MakeFiniteStateMachine(defaultStateName string) *FiniteStateMachine {
	fsm := &FiniteStateMachine{
		state:    defaultStateName,
		handlers: make(map[string]*StateHandler),
	}

	return fsm
}

// AddStateHandler used to add handler for fsm's state
func (fsm *FiniteStateMachine) AddStateHandler(
	stateName string,
	enterHandler func(string),
	runHandler func(),
	exitHandler func(),
) {
	fsm.handlers[stateName] = makeHandler(enterHandler, runHandler, exitHandler)
}

// GoTo used for fsm's state transition
func (fsm *FiniteStateMachine) GoTo(stateName string, base Base) {
	previousState := fsm.state
	previousStateHandler := fsm.getStateHandler()

	if previousStateHandler.exitHandler != nil {
		previousStateHandler.exitHandler()
	}

	fsm.state = stateName

	currentStateHandler := fsm.getStateHandler()

	if currentStateHandler.enterHandler != nil {
		currentStateHandler.enterHandler(previousState)
	}
}

// Equals used to compare fsm's state
func (fsm *FiniteStateMachine) Equals(stateName string) bool {
	return fsm.state == stateName
}

// GetName used to get fsm's state name
func (fsm *FiniteStateMachine) GetName() string {
	return fsm.state
}

// Run used to run fsm's state
func (fsm *FiniteStateMachine) Run(item Modellable) {
	handler := fsm.getStateHandler()

	if handler.runHandler != nil {
		handler.runHandler()
	}
}

// SetTurn used to set new turn
func (fsm *FiniteStateMachine) SetTurn(activeAgent Base, endAt time.Time) {
	fsm.activeAgent = activeAgent
	fsm.currentTurnEndAt = endAt
}

// IsTurnExpired used to check if turn is expired
func (fsm *FiniteStateMachine) IsTurnExpired() bool {
	return time.Now().After(fsm.currentTurnEndAt)
}

// GetEndAt used to get turn's end time
func (fsm *FiniteStateMachine) GetEndAt() int64 {
	return fsm.currentTurnEndAt.Unix()
}

// GetActiveAgent used to get fsm's active agent at current turn
func (fsm *FiniteStateMachine) GetActiveAgent() interface{} {
	return fsm.activeAgent
}

func (fsm *FiniteStateMachine) getStateHandler() *StateHandler {
	return fsm.handlers[fsm.state]
}

// StateHandler used to
type StateHandler struct {
	enterHandler func(string)
	runHandler   func()
	exitHandler  func()
}

func makeHandler(
	enterHandler func(string),
	runHandler func(),
	exitHandler func(),
) *StateHandler {
	handler := &StateHandler{
		enterHandler: enterHandler,
		runHandler:   runHandler,
		exitHandler:  exitHandler,
	}

	return handler
}
