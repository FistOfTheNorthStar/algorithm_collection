package tree

import "strings"

// Product represents a product item
type ProductI struct {
	ProductID int
	Name      string
}

// Node represents a node in the binary tree
type NodeI struct {
	GTIN  string
	Data  *ProductI
	Left  *NodeI
	Right *NodeI
}

// Tree represents a binary search tree
type TreeI struct {
	root *NodeI
}

// NewTree creates a new tree instance
func NewTreeI() *TreeI {
	return &TreeI{}
}

// Insert adds a new node to the tree or updates existing one
func (t *TreeI) InsertI(gtin string, productID int, name string) {
	product := &ProductI{
		ProductID: productID,
		Name:      name,
	}

	newNode := &NodeI{
		GTIN: gtin,
		Data: product,
	}

	if t.root == nil {
		t.root = newNode
		return
	}

	current := t.root
	for {
		if gtin < current.GTIN {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else if gtin > current.GTIN {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		} else {
			// Node exists, update data
			current.Data = product
			return
		}
	}
}

// InOrderTraversal performs an in-order traversal of the tree
func (t *TreeI) InOrderTraversal() string {
	var names []string
	t.inOrderTraversalHelper(t.root, &names)

	if len(names) == 0 {
		return "[]"
	}

	return "[" + strings.Join(names, " -> ") + "]"
}

// inOrderTraversalHelper is a recursive helper for in-order traversal
func (t *TreeI) inOrderTraversalHelper(node *NodeI, names *[]string) {
	if node != nil {
		t.inOrderTraversalHelper(node.Left, names)
		*names = append(*names, node.Data.Name)
		t.inOrderTraversalHelper(node.Right, names)
	}
}

// GetRoot returns the root node (mainly for testing)
func (t *TreeI) GetRoot() *NodeI {
	return t.root
}
