package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenHasSingleElement(t *testing.T) {
	list := New()
	list.PushFront(0)
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
	list.PushBack(nil)
	list.PushBack(nil)
	list.PushBack(nil)
	list.PushFront(nil)
	list.PushFront(nil)
	list.PushFront(nil)

	for counter := 0; uint64(counter) < list.Len(); counter++ {
		assert.Nil(t, list.ElementAt(uint64(counter)).Value())
	}

	assert.Equal(t, uint64(6), list.Len())
	assert.Nil(t, list.Last().Value())
	assert.Nil(t, list.First().Value())
}

func TestWhenRegularChainHasBeenPassed(t *testing.T) {
	list := New()
	list.PushFront(0)
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushBack(-1)
	list.PushBack(-2)
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
	list.PushFront(0)
	assert.Equal(t, uint64(1), list.Len())
	assert.Equal(t, list.First().Value(), list.Last().Value())
}
