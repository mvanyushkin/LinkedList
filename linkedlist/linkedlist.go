package LinkedList

import "errors"

type LinkedList struct {
	len      uint64
	headItem *ListItem
}

func (list *LinkedList) Len() uint64 {
	return list.len
}

func (list *LinkedList) First() *ListItem {
	if list.headItem == nil {
		return nil
	}

	return list.headItem
}

func (list *LinkedList) Last() *ListItem {
	if list.headItem == nil {
		return nil
	}

	return list.headItem.prev
}

func (list *LinkedList) PushFront(value interface{}) {
	list.len++
	newItem := &ListItem{value: value, owner: list}
	newItem.next = list.headItem

	if list.headItem == nil {
		list.headItem = newItem
		// Assign the first element to itself
		list.headItem.prev = list.headItem
	}

	newItem.prev = list.headItem.prev
	list.headItem = newItem
}

func (list *LinkedList) PushBack(value interface{}) {
	list.len++
	newTailItem := &ListItem{value: value, owner: list}
	newTailItem.next = nil

	if list.headItem == nil {
		list.headItem = newTailItem
		list.headItem.prev = newTailItem
	}

	oldTailItem := list.headItem.prev
	oldTailItem.next = newTailItem
	newTailItem.prev = oldTailItem
	list.headItem.prev = newTailItem
}

func (list *LinkedList) Remove(item *ListItem) (error error) {
	if item.owner != list {
		return errors.New("wtf?...")
	}

	right := item.prev
	left := item.next

	right.next = left
	left.prev = right
	item.prev = nil
	item.next = nil
	list.len--
	return nil
}

func (list *LinkedList) FindNode(predicate func(v interface{}) bool) *ListItem {
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

func (list *LinkedList) ElementAt(position uint64) *ListItem {
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

func New() LinkedList {
	return LinkedList{}
}
