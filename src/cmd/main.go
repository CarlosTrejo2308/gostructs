package main

import (
	"fmt"

	"github.com/carlostrejo2308/gostructs/src/list"
)

func main() {
	// Create a linked list
	linkedList := &list.Double{}

	// Add elements
	linkedList.Add(1)
	linkedList.Add(4)
	linkedList.Add("22")
	linkedList.Add(22)
	linkedList.Add(20)

	list.PrintList(linkedList)
	fmt.Println("")

	/*
		// Delete an element
		fmt.Println("Deleting element 4")
		linkedList.Delete(4)
		linkedList.Print()

		// Search an element
		idx, err := linkedList.Search(20)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("element: 20 is in index:", idx)
		}

		// Both traversals
		fmt.Printf("%+v\n", linkedList.Traverse())
		fmt.Printf("%+v\n", linkedList.ReverseTraverse())*/
}
