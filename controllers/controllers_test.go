package controllers

import (
	"testing"

	"github.com/karlseguin/expect"
)

type TesterController struct {
}

// RegisterHooksAndActions used to register before hooks and actions
func (controller *TesterController) RegisterHooksAndActions(hooksHandler *HooksHandler, actionsHandler *ActionsHandler) {
	actionsHandler.RegisterConnectedAction(controller.connected)
	actionsHandler.RegisterDisconnectedAction(controller.connected)
	actionsHandler.RegisterErrorAction(controller.error)
}

func (controller *TesterController) connected(context *Context) (interface{}, error) {
	return nil, nil
}

func (controller *TesterController) disconnected(context *Context) (interface{}, error) {
	return nil, nil
}

func (controller *TesterController) error(context *Context) (interface{}, error) {
	return nil, nil
}

type Tester struct {
	manager *Manager
}

func (tester *Tester) RegistorActions() {
	var err error

	tester.manager.RegisterController("test", &TesterController{})

	_ = err
}

func TestController(t *testing.T) {
	manager := &Manager{}

	manager.Setup()

	tester := &Tester{
		manager: manager,
	}

	expect.Expectify(tester, t)
}
