package main

import (
	"fmt"
	LinkedList "github.com/mvanyushkin/LinkedList/linkedlist"
)

func main() {
	list := LinkedList.New()
	list.PushFront(-1)
	list.PushFront(-2)
	list.PushBack(1)
	list.PushBack(2)

	fmt.Printf(" %v", list.First().Next().Next().Prev().Prev().Value())
	fmt.Printf(" %v", list.Last().Prev().Prev().Next().Next().Value())
}
