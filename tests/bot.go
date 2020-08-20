package tests

import (
	"fmt"

	"github.com/kooinam/fabio/actors"
)

type Bot struct {
	id          string
	updateCount int
	energy      int
}

func (bot *Bot) GetID() string {
	return bot.id
}

func (bot *Bot) RegisterActions(actionsHandler *actors.ActionsHandler) {
	actionsHandler.RegisterAction("Start", bot.start)
	actionsHandler.RegisterAction("Update", bot.update)
	actionsHandler.RegisterAction("Attack", bot.attack)
}

func (bot *Bot) start(context *actors.Context) error {
	var err error

	bot.energy = botEnergyCount

	return err
}

func (bot *Bot) update(context *actors.Context) error {
	var err error

	bot.updateCount++

	return err
}

func (bot *Bot) attack(context *actors.Context) error {
	var err error

	if bot.energy <= 0 {
		err = fmt.Errorf("not enough energy")

		return err
	}

	bot.energy--

	return err
}
