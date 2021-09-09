package list

type List interface {
	Add(interface{}) (int, error)
	Delete(interface{}) error
	Search(interface{}) (int, error)
	//Get(int) interface{}
	//GetSize() int
	//IsEmpty() bool
	//Print()
	//Remove(int) interface{}
	//Set(int, interface{})
}

type Node struct {
	Value interface{}
	Next  *Node
	//Prev  *Node
}
