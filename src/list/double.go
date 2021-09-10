package list

import "fmt"

type Double struct {
	head *DNode
	tail *DNode
	Size int
}

func (d *Double) Add(e interface{}) (int, error) {
	// Return an error if the element is nil
	if e == nil {
		return 0, fmt.Errorf("element is nil")
	}

	// Create the node
	n := &DNode{Value: e, Next: nil, Prev: nil}

	// List is empty
	if d.head == nil {
		d.head = n
		d.tail = n
	} else { // Append the node the the end of the list
		n.Prev = d.tail // Link the the prev link of the node to add
		d.tail.Next = n // Link the next link of the
		d.tail = n
	}
	d.Size++
	return d.Size, nil
}

func (d *Double) Delete(e interface{}) error {
	// Pre: e is not nil, and the list is not empty
	if e == nil {
		return fmt.Errorf("element is nil")
	}
	if d.head == nil {
		return fmt.Errorf("list is empty")
	}

	// Element to remove is the head
	if e == d.head.Value {
		// Only node in the list
		if d.head == d.tail {
			d.head = nil
			d.tail = nil
		} else { // There are more nodes in the list
			d.head = d.head.Next
			d.head.Prev = nil
		}
		d.Size--
		return nil
	}

	// Traverse the list until we find the element are we are at the end of the list
	n := d.head
	for n != nil && e != n.Value {
		n = n.Next
	}

	// Element to remove is the tail
	if n == d.tail {
		d.tail = d.tail.Prev
		d.tail.Next = nil
		d.Size--
		return nil
	} else if n != nil { // Element to remove is in the middle of the list
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
		d.Size--
		return nil
	}

	return fmt.Errorf("element not found")
}

func (d *Double) Traverse() []interface{} {
	n := d.head
	var list []interface{}
	for n != nil {
		list = append(list, n.Value)
		n = n.Next
	}
	return list
}

func (d *Double) ReverseTraverse() []interface{} {
	n := d.tail
	var list []interface{}
	for n != nil {
		list = append(list, n.Value)
		n = n.Prev
	}
	return list
}
