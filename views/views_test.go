package views

import (
	"encoding/json"
	"testing"

	"github.com/kooinam/fabio/models"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/simplerecords"
)

type Player struct {
	simplerecords.Base `json:"-"`
	Name               string `json:"name"`
}

type PlayerView struct {
	*Player
}

func makePlayerView() Viewable {
	playerView := &PlayerView{}

	return playerView
}

func (view *PlayerView) Render(context *Context) interface{} {
	view.Player = context.Item().(*Player)

	return view
}

type Tester struct {
	manager *Manager
}

func (tester *Tester) RegisterViews() {
	tester.manager.RegisterView("player", makePlayerView)

	player1 := &Player{
		Name: "tester1",
	}
	player2 := &Player{
		Name: "player2",
	}
	players := models.MakeList(player1, player2)

	view1 := tester.manager.PrepareRender("player").WithRootKey("player").RenderSingle(player1)
	view2 := tester.manager.PrepareRender("player").RenderSingle(player1)
	views := tester.manager.PrepareRender("player").WithRootKey("players").RenderList(players)

	json1, _ := json.Marshal(view1)
	json2, _ := json.Marshal(view2)
	jsons, _ := json.Marshal(views)

	expect.Expect(string(json1)).To.Equal("{\"player\":{\"name\":\"tester1\"}}")
	expect.Expect(string(json2)).To.Equal("{\"name\":\"tester1\"}")
	expect.Expect(string(jsons)).To.Equal("{\"players\":[{\"name\":\"tester1\"},{\"name\":\"player2\"}]}")
}

func TestView(t *testing.T) {
	manager := &Manager{}

	manager.Setup()

	tester := &Tester{
		manager: manager,
	}

	expect.Expectify(tester, t)
}
