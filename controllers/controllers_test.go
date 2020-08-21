package controllers

import (
	"testing"

	"github.com/karlseguin/expect"
)

type TesterController struct {
}

// RegisterBeforeHooks used to register before action hooks
func (controller *TesterController) RegisterBeforeHooks(hooksHandler *HooksHandler) {
}

// RegisterActions used to add actions
func (controller *TesterController) RegisterActions(actionsHandler *ActionsHandler) {
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
