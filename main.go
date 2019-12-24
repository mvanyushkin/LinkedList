package main

import (
	"fmt"
	LinkedList "github/mvanyushkin/LinkedList/linkedlist"
)

func main() {
	fmt.Println("Creating a new linked list")
	list := LinkedList.New()
	fmt.Printf("Current len is %v \n", list.Len())
	fmt.Println("Adding an item to the list from the head side...")
	list.PushFront(1)
	fmt.Printf("Current len is %v \n", list.Len())

	fmt.Println("Adding an item to the list from the head side...")
	list.PushFront(2)
	fmt.Printf("Current len is %v \n", list.Len())

	fmt.Println("Adding an item to the list from the head side...")
	list.PushFront(3)
	fmt.Printf("Current len is %v \n", list.Len())

	fmt.Println("Adding an item to the list from the back side...")
	list.PushBack(-1)
	fmt.Printf("Current len is %v \n", list.Len())

	fmt.Println("Adding an item to the list from the back side...")
	list.PushBack(-2)
	fmt.Printf("Current len is %v  \n", list.Len())

	fmt.Println("Adding an item to the list from the back side...")
	list.PushBack(-1)
	fmt.Printf("Current len is %v \n", list.Len())

	fmt.Println("The whole list is:")
	for counter := 0; uint64(counter) < list.Len(); counter++ {
		fmt.Printf("%v ", list.ElementAt(uint64(counter)).Value())
	}
}
