package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenHasSingleElement(t *testing.T) {
	list := New()
	list.PushBack(0)
	assert.Equal(t, uint64(1), list.Len())
	assert.Equal(t, list.First().Value(), list.Last().Value())
}

func TestWhenEmpty(t *testing.T) {
	list := New()
	assert.Equal(t, uint64(0), list.Len())
	assert.Nil(t, list.Last())
	assert.Nil(t, list.First())
}

func TestWhenNullsHasBeenPassed(t *testing.T) {
	list := New()
	list.PushFront(nil)
	list.PushFront(nil)
	list.PushFront(nil)
	list.PushBack(nil)
	list.PushBack(nil)
	list.PushBack(nil)

	for counter := 0; uint64(counter) < list.Len(); counter++ {
		assert.Nil(t, list.ElementAt(uint64(counter)).Value())
	}

	assert.Equal(t, uint64(6), list.Len())
	assert.Nil(t, list.Last().Value())
	assert.Nil(t, list.First().Value())
}

func TestWhenRegularChainHasBeenPassed(t *testing.T) {
	list := New()
	list.PushBack(0)
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushFront(-1)
	list.PushFront(-2)
	assert.Equal(t, uint64(6), list.Len())
	assert.Equal(t, -2, list.ElementAt(0).Value())
	assert.Equal(t, -1, list.ElementAt(1).Value())
	assert.Equal(t, 0, list.ElementAt(2).Value())
	assert.Equal(t, 1, list.ElementAt(3).Value())
	assert.Equal(t, 2, list.ElementAt(4).Value())
	assert.Equal(t, 3, list.ElementAt(5).Value())
}

func TestWhenRemovingElement(t *testing.T) {
	list := New()
	list.PushBack(0)
	assert.Equal(t, uint64(1), list.Len())
	assert.Equal(t, list.First().Value(), list.Last().Value())
}

func TestWhenHasSinglePrevAndNextAreNull(t *testing.T) {
	list := New()
	list.PushBack(0)
	assert.Equal(t, uint64(1), list.Len())
	assert.Nil(t, list.First().Prev())
	assert.Nil(t, list.First().Next())
	assert.Nil(t, list.Last().Prev())
	assert.Nil(t, list.Last().Next())

	list = New()
	list.PushFront(0)
	assert.Equal(t, uint64(1), list.Len())
	assert.Nil(t, list.First().Prev())
	assert.Nil(t, list.First().Next())
	assert.Nil(t, list.Last().Prev())
	assert.Nil(t, list.Last().Next())
}

func TestCrushTestSpecialForAlexey(t *testing.T) {
	l := New()
	l.PushFront("C")
	l.PushFront("B")
	l.PushFront("A")
	l.PushBack("X")
	l.PushBack("Y")
	l.PushBack("Z")

	// [A] - [B] - [C] - [X] - [Y] - [Z]

	// Test from the head to the tail

	first := l.First() // [A]

	assert.Nil(t, first.Prev())
	assert.Equal(t, "A", first.Value())
	assert.Equal(t, "B", first.Next().Value())
	assert.Equal(t, "C", first.Next().Next().Value())
	assert.Equal(t, "X", first.Next().Next().Next().Value())
	assert.Equal(t, "Y", first.Next().Next().Next().Next().Value())
	assert.Equal(t, "Z", first.Next().Next().Next().Next().Next().Value())
	assert.Nil(t, first.Next().Next().Next().Next().Next().Next())

	// Test from the tail to the head

	last := l.Last()

	assert.Nil(t, last.Next()) // [Z]
	assert.Equal(t, "Z", last.Value())
	assert.Equal(t, "Y", last.Prev().Value())
	assert.Equal(t, "X", last.Prev().Prev().Value())
	assert.Equal(t, "C", last.Prev().Prev().Prev().Value())
	assert.Equal(t, "B", last.Prev().Prev().Prev().Prev().Value())
	assert.Equal(t, "A", last.Prev().Prev().Prev().Prev().Prev().Value())
	assert.Nil(t, last.Prev().Prev().Prev().Prev().Prev().Prev())

	// Test the middle elem and its own heighbors

	middle := l.ElementAt(2) // [C]

	assert.Equal(t, "C", middle.Value())
	assert.Equal(t, "B", middle.Prev().Value())
	assert.Equal(t, "X", middle.Next().Value())

	// Test cutting the tail element from the list

	l.Remove(l.Last()) // // [A] - [B] - [C] - [X] - [Y] - ....
	last = l.Last()
	assert.Nil(t, last.Next()) // [Y]
	assert.Equal(t, "Y", last.Value())
	assert.Equal(t, "X", last.Prev().Value())

	// Test cutting the head from the list

	l.Remove(l.ElementAt(0)) // .... [B] - [C] - [X] - [Y]

	first = l.First()
	assert.Nil(t, first.Prev())
	assert.Equal(t, "B", first.Value())
	assert.Equal(t, "C", first.Next().Value())

	// Test cutting the middle of the list

	l.Remove(l.ElementAt(1)) //  [B] - ... - [X] - [Y]
	l.Remove(l.ElementAt(1)) //  [B] - ... - [Y]

	first = l.First()
	assert.Nil(t, first.Prev())
	assert.Equal(t, "B", first.Value())
	assert.Equal(t, "Y", first.Next().Value())

	//
	l.Remove(l.ElementAt(1)) //  [B] - ...
	l.Remove(l.ElementAt(0)) //  ...

	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())
	assert.Nil(t, l.headItem)
}

func (list *linkedList) FindNode(predicate func(v interface{}) bool) *ListItem {
	item := list.headItem
	for {
		if predicate(item.value) {
			return item
		}

		if item.next == nil {
			break
		}

		item = item.next
	}

	return nil
}

func (list *linkedList) ElementAt(position uint64) *ListItem {
	item := list.headItem
	for index := 0; ; index++ {
		if uint64(index) == position {
			return item
		}

		if item.next == nil {
			break
		}

		item = item.next
	}

	return nil
}
