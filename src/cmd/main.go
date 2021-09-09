package main

import (
	"fmt"

	"github.com/carlostrejo2308/gostructs/src/list"
)

func main() {
	linkedList := &list.Single{}
	linkedList.Add(1)
	linkedList.Add(4)
	linkedList.Add("22")
	linkedList.Add(22)
	linkedList.Add(20)

	linkedList.Print()
	fmt.Println("")
	linkedList.Delete(4)
	linkedList.Print()

	idx, err := linkedList.Search(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("index:", idx)
	}

	fmt.Printf("%+v", linkedList.ReverseTraverse())
}
