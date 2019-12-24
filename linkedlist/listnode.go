package LinkedList

type ListItem struct {
	prev  *ListItem
	next  *ListItem
	value interface{}
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

func NewItem(value interface{}, prev *ListItem, next *ListItem) *ListItem {
	return &ListItem{value: value, prev: prev, next: next}
}
