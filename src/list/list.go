package list

import "fmt"

type List interface {
	Add(interface{}) (int, error)
	Delete(interface{}) error
	//Search(interface{}) (int, error)
	Traverse() []interface{}
	//Get(int) interface{}
	//GetSize() int
	//IsEmpty() boolPrint()
	//Remove(int) interface{}
	//Set(int, interface{})
}

func PrintList(l List) {
	nodes := l.Traverse()
	for _, node := range nodes {
		fmt.Println(node, ", ")
	}
}

type Node struct {
	Value interface{}
	Next  *Node
}

type DNode struct {
	Value interface{}
	Next  *DNode
	Prev  *DNode
}
