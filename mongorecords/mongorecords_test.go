package mongorecords

import (
	"fmt"
	"testing"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
	"github.com/kooinam/fabio/models"
)

type Task struct {
	Base      `bson:"base,inline"`
	Text      string `bson:"text"`
	Completed bool   `bson:"completed"`
}

func makeTask(collection *models.Collection, hooksHandler *models.HooksHandler) models.Modellable {
	task := &Task{}

	hooksHandler.RegisterInitializeHook(task.Initialize)

	hooksHandler.RegisterValidationHook(task.validateTextLength)

	return task
}

func (task *Task) Initialize(dict *helpers.Dictionary) {
	task.Text = dict.ValueStr("text")
}

func (task *Task) validateTextLength() error {
	var err error

	if len(task.Text) < 10 {
		err = fmt.Errorf("text must be longer than or equal to 10 characters")
	}

	return err
}

type Tester struct {
	manager    *models.Manager
	clientName string
}

func (tester *Tester) Each(f func()) {
	f()

	adapter := tester.manager.Adapter(tester.clientName).(*Adapter)

	for _, collection := range adapter.Collections() {
		ctx := adapter.getTimeoutContext()

		adapter.getCollection(collection.Name()).Drop(ctx)
	}
}

func (tester *Tester) QueryCount() {
	adapter, err := MakeAdapter("mongodb://localhost:27017", "tasker")
	expect.Expect(err).To.Equal(nil)

	tester.manager.RegisterAdapter(tester.clientName, adapter)

	collection := tester.manager.RegisterCollection(tester.clientName, "tasks", makeTask)

	count, err := collection.Query().Count()
	expect.Expect(err).To.Equal(nil)
	expect.Expect(count).To.Equal(int64(0))

	_, err = collection.Create(helpers.H{
		"text": "short",
	})
	expect.Expect(err.Error()).To.Equal("text must be longer than or equal to 10 characters")

	texts := []string{"test1test1", "test2test2", "test3test3"}

	for i, text := range texts {
		item, err := collection.Create(helpers.H{
			"text": text,
		})
		expect.Expect(err).To.Equal(nil)

		task := item.(*Task)
		expect.Expect(task.Text).To.Equal(text)

		count, err = collection.Query().Count()
		expect.Expect(err).To.Equal(nil)
		expect.Expect(count).To.Equal(int64(i + 1))
	}

	i := 0

	err = collection.Query().Each(func(item models.Modellable, err error) bool {
		expect.Expect(err).To.Equal(nil)

		if err == nil {
			task := item.(*Task)
			expect.Expect(task.Text).To.Equal(texts[i])

			count, _ = collection.Query().Where(helpers.H{"_id": task.ID}).Count()
			expect.Expect(count).To.Equal(int64(1))
		}

		if i < len(texts)-1 {
			i++
		}

		return true
	})

	count, _ = collection.Query().Where(helpers.H{"text": "test1test1"}).Count()
	expect.Expect(count).To.Equal(int64(1))

	item, err := collection.Query().First()
	expect.Expect(err).To.Equal(nil)

	if err == nil {
		task := item.(*Task)
		expect.Expect(task.Text).To.Equal("test1test1")

		task.Text = "changed"
		err = task.Save()
		expect.Expect(err.Error()).To.Equal("text must be longer than or equal to 10 characters")

		task.Text = "changedchanged"
		err = task.Save()
		expect.Expect(err).To.Equal(nil)
	}

	item, err = collection.Query().First()
	expect.Expect(err).To.Equal(nil)

	if err == nil {
		task := item.(*Task)
		expect.Expect(task.Text).To.Equal("changedchanged")
	}
}

func TestQuery(t *testing.T) {
	manager := &models.Manager{}
	manager.Setup()

	tester := &Tester{
		manager:    manager,
		clientName: "mongo",
	}

	expect.Expectify(tester, t)
}
