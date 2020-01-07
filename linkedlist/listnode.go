package LinkedList

type ListItem struct {
	prev  *ListItem
	next  *ListItem
	value interface{}
	owner *LinkedList
}

func (item *ListItem) Prev() *ListItem {
	return item.prev
}

func (item *ListItem) Next() *ListItem {
	return item.next
}

func (item *ListItem) Value() interface{} {
	return item.value
}
