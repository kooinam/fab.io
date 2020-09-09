package models

import (
	"time"
)

// FiniteStateMachine used to implement fsm
type FiniteStateMachine struct {
	state            string
	handlers         map[string]*StateHandler
	activeAgent      Modellable
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

// RegisterState used to add handler for fsm's state
func (fsm *FiniteStateMachine) RegisterState(stateName string) *StateHandler {
	stateHandler := makeStateHandler(stateName)
	fsm.handlers[stateName] = stateHandler

	return stateHandler
}

// GoTo used for fsm's state transition
func (fsm *FiniteStateMachine) GoTo(stateName string, item Modellable) {
	previousState := fsm.state
	previousStateHandler := fsm.getStateHandler()

	if previousStateHandler.exitHook != nil {
		previousStateHandler.exitHook()
	}

	fsm.state = stateName

	currentStateHandler := fsm.getStateHandler()

	if currentStateHandler.enterHook != nil {
		currentStateHandler.enterHook(previousState)
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
func (fsm *FiniteStateMachine) Run(item Modellable, dt float64) {
	handler := fsm.getStateHandler()

	if handler.runHook != nil {
		handler.runHook(dt)
	}
}

// SetTurn used to set new turn
func (fsm *FiniteStateMachine) SetTurn(activeAgent Modellable, endAt time.Time) {
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
