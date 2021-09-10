package list

import "fmt"

type Single struct {
	head *Node
	tail *Node
	Size int
}

// Add recives an element and inserts it to the end of the linked list
func (s *Single) Add(element interface{}) error {
	// Return an error if the element is nil
	if element == nil {
		return fmt.Errorf("element is nil")
	}
	// Creante a node
	n := &Node{element, nil}

	// If the list is empty, insert it there
	if s.head == nil {
		s.head = n
		s.tail = n
	} else { // Insert it at the end of the list
		s.tail.Next = n
		s.tail = n
	}

	s.Size++
	return nil
}

// Search will search an element in the linked list and returns the index of
// the first match it finds. If the list is empty or the element is not in the list,
// it will return an error
func (s *Single) Search(element interface{}) (int, error) {
	idx := 0
	// The list is empty
	if s.head == nil {
		return 0, fmt.Errorf("list is empty")
	}
	// Traverse all the list
	for n := s.head; n != nil; n = n.Next {
		// If the element is found, return the index
		if n.Value == element {
			return idx, nil
		}
		idx++
	}

	// We finish traversing and the element was not found
	return 0, fmt.Errorf("item not found")
}

// Delete will delete the first node it maches
// It will return an error if the list is empty, or
// if the element was not in the list
func (s *Single) Delete(element interface{}) error {
	// The list is empty
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	// The element is in the head
	if s.head.Value == element {
		s.head = s.head.Next
		s.Size--
		return nil
	}

	// Start searching on node 1
	n := s.head

	// Search on all the list
	for n.Next != nil {
		// Next node is the one to delete
		if n.Next.Value == element {
			n.Next = n.Next.Next
			s.Size--
			return nil
		}
		n = n.Next
	}
	return fmt.Errorf("item not found")
}

// Traverse returns an array of all the elements in
// the linked list
func (s *Single) Traverse() []interface{} {
	var result []interface{}
	if s.head == nil {
		return result
	}
	for n := s.head; n != nil; n = n.Next {
		result = append(result, n.Value)
	}
	return result
}

// ReverseTraverse returns an array of the linked
// list in a reversed order.
func (s *Single) ReverseTraverse() []interface{} {
	var result []interface{}

	// If there's an element in the tail, aka list not empty
	if s.tail != nil {
		curr := s.tail

		for curr != s.head {
			prev := s.head // Start searching from the head

			// Traverse the linkedlist, until the next one is the one we are going to add
			for prev.Next != curr {
				prev = prev.Next
			}

			// Add the value of the one we were going to add, and update the current
			// one with his previous node
			result = append(result, curr.Value)
			curr = prev
		}
		// Add the head
		result = append(result, curr.Value)
	}
	return result
}
