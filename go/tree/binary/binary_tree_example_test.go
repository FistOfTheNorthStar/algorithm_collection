package tree

import "testing"

func TestProductTree(t *testing.T) {
	// Create test products
	products := []*Product{
		{ProductID: 1, Name: "Product 1", Price: 10.0, ManufacturerName: "Manufacturer 1"},
		{ProductID: 2, Name: "Product 2", Price: 20.0, ManufacturerName: "Manufacturer 2"},
		{ProductID: 3, Name: "Product 3", Price: 30.0, ManufacturerName: "Manufacturer 3"},
		{ProductID: 4, Name: "Product 4", Price: 40.0, ManufacturerName: "Manufacturer 4"},
		{ProductID: 5, Name: "Product 5", Price: 50.0, ManufacturerName: "Manufacturer 5"},
		{ProductID: 6, Name: "Product 6", Price: 60.0, ManufacturerName: "Manufacturer 6"},
		{ProductID: 7, Name: "Product 7", Price: 70.0, ManufacturerName: "Manufacturer 7"},
		{ProductID: 8, Name: "Product 8", Price: 80.0, ManufacturerName: "Manufacturer 8"},
		{ProductID: 9, Name: "Product 9", Price: 90.0, ManufacturerName: "Manufacturer 9"},
	}

	t.Run("find node operations", func(t *testing.T) {
		tree := NewTree()

		// Insert test data
		tree.Insert("04000345706564", []*Product{products[0]})
		tree.Insert("07611400983416", []*Product{products[1]})
		tree.Insert("07611400989104", []*Product{products[2], products[3]})
		tree.Insert("07611400989111", []*Product{products[4]})
		tree.Insert("07611400990292", []*Product{products[5], products[6], products[7]})

		// Test finding non-existent node
		if node := tree.Find("07611400983324"); node != nil {
			t.Error("Expected nil for non-existent GTIN")
		}

		// Insert and find new node
		tree.Insert("07611400983324", []*Product{products[8]})
		node := tree.Find("07611400983324")

		if node == nil {
			t.Error("Expected to find inserted node")
		}

		if node != nil && node.GTIN != "07611400983324" {
			t.Errorf("Wrong GTIN found. Got %s, want 07611400983324", node.GTIN)
		}
	})
}
