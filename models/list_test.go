package models

import (
	"testing"

	"github.com/karlseguin/expect"
)

type Item struct {
	Name string
}

func (item *Item) InitializeBase(*Context)        {}
func (item *Item) GetID() string                  { return "" }
func (item *Item) Save() error                    { return nil }
func (item *Item) GetHooksHandler() *HooksHandler { return nil }
func (item *Item) Store()                         {}
func (item *Item) StoreInList(*List)              {}

type ListTester struct{}

func (tester *ListTester) Sort() {
	list := MakeList(
		&Item{
			Name: "123",
		},
		&Item{
			Name: "cde",
		},
		&Item{
			Name: "abc",
		},
		&Item{
			Name: "aac",
		},
		&Item{
			Name: "ccc",
		},
		&Item{
			Name: "b",
		},
	)

	sorted := list.Sort(func(e1 Modellable, e2 Modellable) bool {
		item1 := e1.(*Item)
		item2 := e2.(*Item)

		return item1.Name < item2.Name
	})

	items := sorted.Items()
	expect.Expect(items[0].(*Item).Name).ToEqual("123")
	expect.Expect(items[1].(*Item).Name).ToEqual("aac")
	expect.Expect(items[2].(*Item).Name).ToEqual("abc")
	expect.Expect(items[3].(*Item).Name).ToEqual("b")
	expect.Expect(items[4].(*Item).Name).ToEqual("ccc")
	expect.Expect(items[5].(*Item).Name).ToEqual("cde")
}

func TestList(t *testing.T) {
	tester := &ListTester{}

	expect.Expectify(tester, t)
}
