package LinkedList

import "errors"

type linkedList struct {
	len      uint64
	headItem *ListItem
}

func (list *linkedList) Len() uint64 {
	return list.len
}

func (list *linkedList) First() *ListItem {
	if list.headItem == nil {
		return nil
	}

	return list.headItem
}

func (list *linkedList) Last() *ListItem {
	if list.headItem == nil {
		return nil
	}

	return list.headItem.prev
}

func (list *linkedList) PushFront(value interface{}) {
	list.len++
	newItem := &ListItem{value: value, owner: list}
	if list.headItem == nil {
		list.headItem = newItem
		list.headItem.next = newItem
		list.headItem.prev = newItem
	} else {
		newItem.next = list.headItem
		newItem.prev = list.headItem.prev
		list.headItem.prev.next = newItem
		list.headItem.prev = newItem
		list.headItem = newItem
	}
}

func (list *linkedList) PushBack(value interface{}) {
	list.len++
	newItem := &ListItem{value: value, owner: list}
	if list.headItem == nil {
		list.headItem = newItem
		list.headItem.next = newItem
		list.headItem.prev = newItem
	} else {
		newItem.next = list.headItem
		newItem.prev = list.headItem.prev
		list.headItem.prev.next = newItem
		list.headItem.prev = newItem
	}
}

func (list *linkedList) Remove(item *ListItem) (error error) {
	if item == nil || item.owner != list {
		return errors.New("wtf?...")
	}

	left := item.prev
	right := item.next

	left.next = right
	right.prev = left

	if item == list.headItem {
		if right == list.headItem {
			list.headItem = nil
		} else {
			list.headItem = right
		}
	}

	item.prev = nil
	item.next = nil
	item.owner = nil
	list.len--
	return nil
}

func New() linkedList {
	return linkedList{}
}
