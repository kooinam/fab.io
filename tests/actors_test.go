package tests

import (
	"testing"
	"time"

	"github.com/karlseguin/expect"
	fab "github.com/kooinam/fabio"
	"github.com/kooinam/fabio/helpers"
)

var botEnergyCount = 0

type ActorTester struct {
}

func (tester *ActorTester) RegistorActions() {
	var err error
	botEnergyCount = 5
	fab.Setup()

	bot := &Bot{}

	actor := fab.ActorManager().RegisterActor("bot", bot)

	updateCount := 2

	time.Sleep(time.Duration(updateCount+1) * time.Second)

	expect.Expect(bot.energy).To.Equal(5)
	expect.Expect(bot.updateCount).To.Equal(updateCount)

	err = fab.ActorManager().Request(actor.Identifier(), "attack", helpers.H{})
	expect.Expect(err.Error()).To.Equal("no action found")

	err = fab.ActorManager().Request(actor.Identifier(), "Attack", helpers.H{})
	expect.Expect(bot.energy).To.Equal(4)
}

func TestActorRegisterActions(t *testing.T) {
	expect.Expectify(new(ActorTester), t)
}
