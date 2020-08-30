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

	counter := collection.Query().Count()
	expect.Expect(counter.Status()).To.Equal(models.StatusSuccess)
	expect.Expect(counter.Count()).To.Equal(int64(0))

	result := collection.Create(helpers.H{
		"text": "short",
	})
	expect.Expect(result.ErrorMessage()).To.Equal("text must be longer than or equal to 10 characters")

	texts := []string{"test1test1", "test2test2", "test3test3"}

	for i, text := range texts {
		result = collection.Create(helpers.H{
			"text": text,
		})
		expect.Expect(result.Status()).To.Equal(models.StatusSuccess)

		task := result.Item().(*Task)
		expect.Expect(task.Text).To.Equal(text)

		counter = collection.Query().Count()
		expect.Expect(counter.Status()).To.Equal(models.StatusSuccess)
		expect.Expect(counter.Count()).To.Equal(int64(i + 1))
	}

	i := 0

	err = collection.Query().Each(func(item models.Modellable, err error) bool {
		expect.Expect(err).To.Equal(nil)

		if err == nil {
			task := item.(*Task)
			expect.Expect(task.Text).To.Equal(texts[i])

			counter = collection.Query().Where(helpers.H{"_id": task.ID}).Count()
			expect.Expect(counter.Count()).To.Equal(int64(1))
		}

		if i < len(texts)-1 {
			i++
		}

		return true
	})

	counter = collection.Query().Where(helpers.H{"text": "test1test1"}).Count()
	expect.Expect(counter.Count()).To.Equal(int64(1))

	result = collection.Query().First()
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)

	if result.StatusSuccess() {
		task := result.Item().(*Task)
		expect.Expect(task.Text).To.Equal("test1test1")

		task.Text = "changed"
		err = task.Save()
		expect.Expect(err.Error()).To.Equal("text must be longer than or equal to 10 characters")

		task.Text = "changedchanged"
		err = task.Save()
		expect.Expect(err).To.Equal(nil)
	}

	result = collection.Query().First()
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)

	if result.StatusSuccess() {
		task := result.Item().(*Task)
		expect.Expect(task.Text).To.Equal("changedchanged")
	}

	result = collection.Query().Where(helpers.H{
		"text": "no",
	}).First()
	expect.Expect(result.Status()).To.Equal(models.StatusNotFound)

	result = collection.Query().Where(helpers.H{
		"text": "no",
	}).FirstOrCreate(helpers.H{})
	expect.Expect(result.Status()).To.Equal(models.StatusError)
	expect.Expect(result.ErrorMessage()).To.Equal("text must be longer than or equal to 10 characters")

	result = collection.Query().Where(helpers.H{
		"text": "nooooooooooook",
	}).FirstOrCreate(helpers.H{})
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)

	id := result.Item().GetID()
	result = collection.Query().Where(helpers.H{
		"text": "nooooooooooook",
	}).FirstOrCreate(helpers.H{})
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)
	expect.Expect(result.Item().GetID()).To.Equal(id)

	result = collection.Query().Find(id)
	expect.Expect(result.Status()).To.Equal(models.StatusSuccess)
	expect.Expect(result.Item().GetID()).To.Equal(id)
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
