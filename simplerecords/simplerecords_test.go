package simplerecords

import (
	"testing"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

type Task struct {
	Base
	Text      string
	Completed bool
}

func makeTask(collection *models.Collection, hooksHandler *models.HooksHandler) models.Modellable {
	task := &Task{}

	hooksHandler.RegisterInitializeHook(task.Initialize)

	return task
}

func (task *Task) Initialize(dict *helpers.Dictionary) {
	task.Text = dict.ValueStr("text")
}

type Tester struct {
	clientName string
	manager    *models.Manager
}

func (tester *Tester) QueryCount() {
	adapter := MakeAdapter()

	tester.manager.RegisterAdapter(tester.clientName, adapter)

	collection := tester.manager.RegisterCollection(tester.clientName, "tasks", makeTask)

	item, err := collection.Create(helpers.H{
		"text": "test",
	})

	expect.Expect(err).To.Equal(nil)
	expect.Expect(item.GetID()).To.Equal("1")
}

func TestQuery(t *testing.T) {
	manager := &models.Manager{}
	manager.Setup()

	tester := &Tester{
		clientName: "simple",
		manager:    manager,
	}

	expect.Expectify(tester, t)
}
