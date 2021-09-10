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
