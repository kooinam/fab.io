package actors

import (
	"fmt"
	"testing"
	"time"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fab.io/helpers"
)

type ChildBot struct {
	id      string
	tester  *Tester
	started bool
	updated bool
}

func (bot *ChildBot) GetActorIdentifier() string {
	return bot.id
}

func (bot *ChildBot) RegisterActorActions(actionsHandler *ActionsHandler) {
	actionsHandler.RegisterAction("Start", bot.start)
	actionsHandler.RegisterAction("Update", bot.update)
	// actionsHandler.RegisterAction("Attack", bot.attack)
}

func (bot *ChildBot) start(context *Context) error {
	var err error

	bot.started = true

	return err
}

func (bot *ChildBot) update(context *Context) error {
	var err error

	bot.updated = true

	return err
}

type TesterBot struct {
	id                string
	tester            *Tester
	timer             float64
	replenishInterval float64
	energy            int
	children          []*ChildBot
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
	bot.children = []*ChildBot{}

	childIDs := []string{"child-1", "child-2"}

	for _, childID := range childIDs {
		childBot := &ChildBot{
			tester: bot.tester,
			id:     childID,
		}
		bot.children = append(bot.children, childBot)

		err = bot.tester.manager.RegisterChildActor(bot, childBot)
		expect.Expect(err).To.Equal(nil)
	}

	childBot := &ChildBot{
		tester: bot.tester,
		id:     "child-1",
	}

	err = bot.tester.manager.RegisterChildActor(bot, childBot)
	expect.Expect(err.Error()).To.Equal("actor already registered")

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

func (tester *Tester) Each(f func()) {
	manager := &Manager{}
	manager.Setup()

	tester.manager = manager

	f()
}

func (tester *Tester) RegistorActions() {
	var err error

	bot := &TesterBot{
		tester: tester,
		id:     "bot-1",
	}

	err = tester.manager.RegisterActor(bot)
	expect.Expect(err).To.Equal(nil)

	err = tester.manager.RegisterActor(bot)
	expect.Expect(err.Error()).To.Equal("actor already registered")

	time.Sleep(1 * time.Second)

	expect.Expect(bot.energy).To.Equal(0)
	expect.Expect(bot.replenishInterval).To.Equal(float64(2))

	time.Sleep(4 * time.Second)

	expect.Expect(bot.energy).To.Equal(2)

	err = tester.manager.Request(bot.GetActorIdentifier(), "attack", helpers.H{})
	expect.Expect(err.Error()).To.Equal("no action found")

	err = tester.manager.Request(bot.GetActorIdentifier(), "Attack", helpers.H{})
	expect.Expect(bot.energy).To.Equal(1)
}

func (tester *Tester) SpawnChildActors() {
	var err error

	bot := &TesterBot{
		tester: tester,
		id:     "bot-1",
	}

	err = tester.manager.RegisterActor(bot)
	expect.Expect(err).To.Equal(nil)

	time.Sleep(1 * time.Second)

	expect.Expect(len(bot.children)).To.Equal(2)

	time.Sleep(2 * time.Second)

	for _, child := range bot.children {
		expect.Expect(child.started).To.Equal(true)
		expect.Expect(child.updated).To.Equal(true)
	}
}

func TestActor(t *testing.T) {
	tester := &Tester{}

	expect.Expectify(tester, t)
}
