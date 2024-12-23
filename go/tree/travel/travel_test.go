package tree

import "testing"

func TestTreeTraversal(t *testing.T) {
	tree := NewTreeI()

	// Insert test data
	tree.InsertI("0510414", 12, "apple")
	tree.InsertI("0301712", 14, "cherry")
	tree.InsertI("0814101", 13, "mango")
	tree.InsertI("0401312", 19, "orange")
	tree.InsertI("0911741", 11, "banana")
	tree.InsertI("0617411", 16, "grape")
	tree.InsertI("0210417", 18, "kiwi")

	expected := "[kiwi -> cherry -> orange -> apple -> grape -> mango -> banana]"
	result := tree.InOrderTraversal()

	if result != expected {
		t.Errorf("InOrderTraversal() = %q; want %q", result, expected)
	}
}

func TestEmptyTree(t *testing.T) {
	tree := NewTreeI()
	result := tree.InOrderTraversal()
	expected := "[]"

	if result != expected {
		t.Errorf("InOrderTraversal() = %q; want %q", result, expected)
	}
}

func TestUpdateExisting(t *testing.T) {
	tree := NewTreeI()

	// Insert initial node
	tree.InsertI("0510414", 12, "apple")

	// Update the same node
	tree.InsertI("0510414", 13, "new apple")

	expected := "[new apple]"
	result := tree.InOrderTraversal()

	if result != expected {
		t.Errorf("InOrderTraversal() = %q; want %q", result, expected)
	}
}
