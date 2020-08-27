package models

import (
	"fmt"
	"testing"

	"github.com/karlseguin/expect"
	"github.com/kooinam/fabio/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	Base      `bson:"base,inline"`
	Text      string `bson:"text"`
	Completed bool   `bson:"completed"`
}

func makeTask(collection *Collection, hooksHandler *HooksHandler) Modellable {
	task := &Task{}

	hooksHandler.RegisterAfterInitializeHook(task.afterInitialize)

	hooksHandler.RegisterValidationHook(task.validateTextLength)

	return task
}

func (task *Task) afterInitialize(dict *helpers.Dictionary) {
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
	manager *Manager
}

func (tester *Tester) Each(f func()) {
	f()

	for _, collection := range tester.manager.collections {
		ctx := tester.manager.adapter.getTimeoutContext()
		tester.manager.adapter.getCollection(collection.name).Drop(ctx)
	}
}

func (tester *Tester) QueryCount() {
	err := tester.manager.RegisterAdapter("mongodb://localhost:27017", "tasker")
	expect.Expect(err).To.Equal(nil)

	collection := tester.manager.CreateCollection("tasks", makeTask)

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

	err = collection.Query().Each(func(item Modellable, err error) {
		expect.Expect(err).To.Equal(nil)

		if err == nil {
			task := item.(*Task)
			expect.Expect(task.Text).To.Equal(texts[i])

			count, _ = collection.Query().Where(bson.M{"_id": task.ID}).Count()
			expect.Expect(count).To.Equal(int64(1))
		}

		if i < len(texts)-1 {
			i++
		}
	})

	count, _ = collection.Query().Where(bson.M{"text": "test1test1"}).Count()
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
	manager := &Manager{}

	manager.Setup()

	tester := &Tester{
		manager: manager,
	}

	expect.Expectify(tester, t)
}
