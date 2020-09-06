package actors

import (
	"fmt"
	"testing"
	"time"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
)

type TesterBot struct {
	id                string
	tester            *Tester
	timer             float64
	replenishInterval float64
	energy            int
}

func (bot *TesterBot) GetActorIdentifier() string {
	return bot.id
}

func (bot *TesterBot) RegisterActorActions(actionsHandler *ActionsHandler) {
	actionsHandler.RegisterAction("Start", bot.start)
	actionsHandler.RegisterAction("Update", bot.update)
	actionsHandler.RegisterAction("Attack", bot.attack)
}

func (bot *TesterBot) start(context *Context) error {
	var err error

	bot.replenishInterval = 2

	return err
}

func (bot *TesterBot) update(context *Context) error {
	var err error

	dt := context.ParamsFloat64("dt", 0)
	bot.timer += dt

	if bot.timer >= bot.replenishInterval {
		bot.timer -= bot.replenishInterval

		bot.energy++
	}

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
	manager *Manager
}

func (tester *Tester) RegistorActions() {
	var err error

	bot := &TesterBot{
		tester: tester,
	}

	actor := tester.manager.RegisterActor(bot)

	time.Sleep(1 * time.Second)

	expect.Expect(bot.energy).To.Equal(0)
	expect.Expect(bot.replenishInterval).To.Equal(float64(2))

	time.Sleep(10 * time.Second)

	expect.Expect(bot.energy).To.Equal(5)

	err = tester.manager.Request(actor.Identifier(), "attack", helpers.H{})
	expect.Expect(err.Error()).To.Equal("no action found")

	err = tester.manager.Request(actor.Identifier(), "Attack", helpers.H{})
	expect.Expect(bot.energy).To.Equal(4)
}

func TestActor(t *testing.T) {
	manager := &Manager{}

	manager.Setup()

	tester := &Tester{
		manager: manager,
	}

	expect.Expectify(tester, t)
}
