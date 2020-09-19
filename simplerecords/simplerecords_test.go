package simplerecords

import (
	"testing"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

type Person struct {
	Base
	Name string
}

func makePerson(context *models.Context) {
	person := &Person{}
	context.SetItem(person)

	context.HooksHandler().RegisterInitializeHook(person.initialize)
}

func (person *Person) initialize(attributes *helpers.Dictionary) {
	person.Name = attributes.ValueStr("name")
}

type Task struct {
	Base
	Text      string
	Completed bool
	Owner     *models.BelongsTo
}

func makeTask(context *models.Context) {
	task := &Task{}
	context.SetItem(task)

	collection := context.Attributes().Value("collection").(*models.Collection)
	task.Owner = context.AssociationsHandler().RegisterBelongsTo(collection)

	context.HooksHandler().RegisterInitializeHook(task.initialize)
}

func (task *Task) initialize(attributes *helpers.Dictionary) {
	task.Text = attributes.ValueStr("text")
	task.Owner.SetKey(attributes.ValueStr("ownerID"))
}

type Tester struct {
	clientName string
	manager    *models.Manager
}

func (tester *Tester) QueryCount() {
	adapter := MakeAdapter()

	tester.manager.RegisterAdapter(tester.clientName, adapter)

	collection1 := tester.manager.RegisterCollection(tester.clientName, "people", makePerson)
	result := collection1.CreateWithOptions(helpers.H{
		"name": "tester1",
	}, models.Options().WithShouldStore(true))
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)
	expect.Expect(result.Item().GetID()).To.Equal("1")

	collection2 := tester.manager.RegisterCollection(tester.clientName, "tasks", makeTask)
	result = collection2.CreateWithOptions(helpers.H{
		"text":       "test",
		"collection": collection1,
		"ownerID":    "1",
	}, models.Options().WithShouldStore(true))
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)
	expect.Expect(result.Item().GetID()).To.Equal("2")

	task := result.Item().(*Task)
	person := task.Owner.Item().(*Person)
	expect.Expect(person.Name).To.Equal("tester1")

	result = collection1.CreateWithOptions(helpers.H{
		"name": "tester2",
	}, models.Options().WithShouldStore(true))
	result = collection1.CreateWithOptions(helpers.H{
		"name": "tester3",
	}, models.Options().WithShouldStore(true))
	result = collection1.Query().Where(helpers.H{
		"Name": "tester2",
	}).First()
	found := result.Item().(*Person)
	expect.Expect(collection1.Query().Count().Count()).To.Equal(int64(3))
	expect.Expect(result.StatusSuccess()).To.Equal(true)
	expect.Expect(found.Name).To.Equal("tester2")
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
