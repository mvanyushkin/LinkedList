package LinkedList

type LinkedList struct {
	len      uint64
	tailItem *ListItem
	headItem *ListItem
}

func (list *LinkedList) Len() uint64 {
	return list.len
}

func (list *LinkedList) First() *ListItem {
	return list.tailItem
}

func (list *LinkedList) Last() *ListItem {
	return list.headItem
}

func (list *LinkedList) PushFront(value interface{}) {
	list.len++
	newItem := NewItem(value, nil, nil)
	wasInitialized := list.tryInitialize(newItem)
	if wasInitialized {
		return
	}

	list.headItem.next = newItem
	newItem.prev = list.headItem
	list.headItem = newItem
}

func (list *LinkedList) PushBack(value interface{}) {
	list.len++
	newItem := NewItem(value, nil, nil)
	wasInitialized := list.tryInitialize(newItem)
	if wasInitialized {
		return
	}

	list.tailItem.prev = newItem
	newItem.next = list.tailItem
	list.tailItem = newItem
}

func (list *LinkedList) Remove(item *ListItem) {
	if item == list.tailItem {
		list.tailItem = item.next
		list.tailItem.prev = nil
		return
	}

	if item == list.headItem {
		list.headItem = item.prev
		list.headItem.next = nil
		return
	}

	item.prev.next = item.next
	item.next.prev = item.prev
}

func (list *LinkedList) FindNode(predicate func(v interface{}) bool) *ListItem {
	currentItem := list.tailItem
	for index := 0; currentItem.next != nil; index++ {
		if uint64(index) > list.len+1 {
			panic("Houston, we got a problem! Eternal iteration has been detected!")
		}

		if predicate(currentItem) {
			return currentItem
		}

		currentItem = currentItem.next
	}
	return nil
}

func (list *LinkedList) ElementAt(position uint64) *ListItem {
	var currentItem *ListItem
	for index := 0; uint64(index) < position+1; index++ {
		if currentItem == nil {
			currentItem = list.tailItem
		} else {
			currentItem = currentItem.next
		}

	}
	return currentItem
}

func (list *LinkedList) tryInitialize(newItem *ListItem) bool {
	if list.tailItem == nil && list.headItem == nil {
		newItem.next = newItem
		newItem.prev = newItem
		list.tailItem = newItem
		list.headItem = newItem
		return true
	}
	return false
}

func New() LinkedList {
	return LinkedList{}
}
