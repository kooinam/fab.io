package views

import (
	"encoding/json"
	"testing"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
)

type Player struct {
	Name string "json:\"name\""
}

type PlayerView struct {
	*Player
}

func makePlayerView() Viewable {
	playerView := &PlayerView{}

	return playerView
}

func (view *PlayerView) Render(properties *helpers.Dictionary) interface{} {
	view.Player = properties.Value("player").(*Player)

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
	players := []interface{}{
		player1,
		player2,
	}

	view1 := tester.manager.RenderView("player", helpers.H{
		"player": player1,
	}, "player")
	view2 := tester.manager.RenderViewWithoutRootKey("player", helpers.H{
		"player": player1,
	})
	views := tester.manager.RenderViews("player", helpers.MakeHashes(players, func(item interface{}) helpers.H {
		return helpers.H{
			"player": item,
		}
	}), "players")

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
