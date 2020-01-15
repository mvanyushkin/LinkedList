package LinkedList

type ListItem struct {
	prev  *ListItem
	next  *ListItem
	value interface{}
	owner *linkedList
}

func (item *ListItem) Prev() *ListItem {
	if item.prev != nil && item.owner.headItem != item {
		return item.prev
	}

	return nil
}

func (item *ListItem) Next() *ListItem {
	if item.next != nil && item.owner.headItem.prev != item {
		return item.next
	}

	return nil
}

func (item *ListItem) Value() interface{} {
	return item.value
}
