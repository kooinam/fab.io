package helpers

import (
	"testing"

	"github.com/karlseguin/expect"
)

type Player struct {
	Name string
}

type Tester struct{}

func (tester *Tester) RegistorActions() {
	player := &Player{
		Name: "test",
	}

	expect.Expect(GetFieldValueByName(player, "Name")).To.Equal("test")
	expect.Expect(GetFieldValueByName(player, "name")).To.Equal(nil)
}

func TestActor(t *testing.T) {
	tester := &Tester{}

	expect.Expectify(tester, t)
}
