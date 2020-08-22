package actors

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
)

type TesterBot struct {
	id          string
	tester      *Tester
	updateCount int
	energy      int
}

func (bot *TesterBot) GetID() string {
	return bot.id
}

func (bot *TesterBot) RegisterActions(actionsHandler *ActionsHandler) {
	actionsHandler.RegisterAction("Start", bot.start)
	actionsHandler.RegisterAction("Update", bot.update)
	actionsHandler.RegisterAction("Attack", bot.attack)
}

func (bot *TesterBot) start(context *Context) error {
	var err error

	bot.energy = bot.tester.botEnergyCount

	return err
}

func (bot *TesterBot) update(context *Context) error {
	var err error

	bot.updateCount++

	return err
}

func (bot *TesterBot) attack(context *Context) error {
	var err error

	if bot.energy <= 0 {
		err = fmt.Errorf("not enough energy")

		return err
	}

	bot.energy--

	return err
}

type Tester struct {
	manager        *Manager
	botEnergyCount int
}

func (tester *Tester) RegistorActions() {
	var err error

	bot := &TesterBot{
		tester: tester,
	}

	actor := tester.manager.RegisterActor("bot", bot)

	updateCount := 2

	time.Sleep(time.Duration(updateCount+1) * time.Second)

	expect.Expect(bot.energy).To.Equal(5)
	expect.Expect(bot.updateCount).To.Equal(updateCount)

	err = tester.manager.Request(actor.Identifier(), "attack", helpers.H{})
	expect.Expect(err.Error()).To.Equal("no action found")

	err = tester.manager.Request(actor.Identifier(), "Attack", helpers.H{})
	expect.Expect(bot.energy).To.Equal(4)

	var params helpers.H
	json.Unmarshal([]byte(""), &params)
}

func TestActor(t *testing.T) {
	manager := &Manager{}

	manager.Setup()

	tester := &Tester{
		manager:        manager,
		botEnergyCount: 5,
	}

	expect.Expectify(tester, t)
}
