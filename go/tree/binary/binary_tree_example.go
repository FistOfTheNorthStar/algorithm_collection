package tree

// Product represents a product item
type Product struct {
	ProductID        int
	Name             string
	Price            float64
	ManufacturerName string
}

// Node represents a node in the binary tree
type Node struct {
	GTIN  string
	Data  []*Product
	Left  *Node
	Right *Node
}

// Tree represents a binary search tree of products
type Tree struct {
	root *Node
}

// NewNode creates a new tree node
func NewNode(gtin string, data []*Product) *Node {
	return &Node{
		GTIN: gtin,
		Data: data,
	}
}

// NewTree creates a new product tree
func NewTree() *Tree {
	return &Tree{}
}

// Insert adds a new node to the tree
func (t *Tree) Insert(gtin string, data []*Product) {
	newNode := NewNode(gtin, data)

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
			// Node already exists
			return
		}
	}
}

// Find searches for a node by GTIN
func (t *Tree) Find(gtin string) *Node {
	current := t.root
	if current == nil {
		return nil
	}

	for current != nil && current.GTIN != gtin {
		if gtin < current.GTIN {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return current
}
