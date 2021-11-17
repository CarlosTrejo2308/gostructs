package tree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{}
}

func (t *Tree) Insert(val int) {
	// Tree is empty
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else { // Insrt the element to the existing tree
		t.Root.insertNode(val)
	}

}

func (current *TreeNode) insertNode(value int) {
	// Inser to the left
	if value < current.Val {
		if current.Left == nil {
			current.Left = &TreeNode{Val: value}
		} else {
			current.Left.insertNode(value)
		}
	} else { // Inser to the right
		if current.Right == nil {
			current.Right = &TreeNode{Val: value}
		} else {
			current.Right.insertNode(value)
		}
	}
}

func (t *Tree) Contains(val int) bool {
	// Tree is empty
	if t.Root == nil {
		return false
	}
	return t.Root.searchNode(val)
}

func (current *TreeNode) searchNode(value int) bool {
	// Node is empty
	if current == nil {
		return false
	}
	// Node is found
	if current.Val == value {
		return true
	}

	// Search left
	if value < current.Val {
		return current.Left.searchNode(value)
	} else { // Search right
		return current.Right.searchNode(value)
	}
}

func (t *Tree) Delete(val int) bool {

	nodeToRemove := t.Root.FindNode(val)

	// Node is not found
	if nodeToRemove == nil {
		return false
	}

	parent := t.Root.FindParent(val)

	// Node is root
	if parent == nil {
		t.Root = nil
		return true
	}

	count := t.Root.Size()

	// Removing the only node in the tree
	if count == 1 {
		t.Root = nil
		return true
	}

	// Case 1: Node is a leaf
	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		// Node is a left child
		if nodeToRemove.Val < parent.Val {
			parent.Left = nil

			// Node is a right child
		} else {
			parent.Right = nil
		}
	}

	// Case 2 and 3: Node has one child
	if nodeToRemove.Left == nil || nodeToRemove.Right == nil {
		// Case 2: Node is a left child
		if nodeToRemove.Val < parent.Val {
			if nodeToRemove.Left != nil {
				parent.Left = nodeToRemove.Left
			} else {
				parent.Left = nodeToRemove.Right
			}

			// Case 3: Node is a right child
		} else {
			if nodeToRemove.Left != nil {
				parent.Right = nodeToRemove.Left
			} else {
				parent.Right = nodeToRemove.Right
			}
		}
	}

}
