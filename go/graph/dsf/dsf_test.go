package graph

import (
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex("Berlin")

	if len(g.GetMapOfVertex()) != 1 {
		t.Errorf("Expected map size of 1, got %d", len(g.GetMapOfVertex()))
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	city1, city2 := "Berlin", "Leipzig"

	g.AddVertex(city1) // location 0
	g.AddVertex(city2) // location 1
	g.AddEdge(city1, city2)

	matrix := g.GetMatrixOfAdjVertex()
	if matrix[0][1] != 1 || matrix[1][0] != 1 {
		t.Error("Edge was not properly added")
	}
}

func TestDFS(t *testing.T) {
	g := NewGraph()
	cities := []string{
		"Berlin", "Leipzig", "Dresden", "Nürnberg",
		"Hannover", "Rostock", "Dortmund", "Frankfurt",
		"Stuttgart", "München", "Magdeburg", "Bremen",
	}

	// Add vertices
	for _, city := range cities {
		g.AddVertex(city)
	}

	// Add edges
	edges := [][2]string{
		{"Berlin", "Leipzig"}, {"Leipzig", "Dresden"},
		{"Dresden", "Nürnberg"}, {"Nürnberg", "München"},
		{"Magdeburg", "Hannover"}, {"Hannover", "Dortmund"},
		{"Dortmund", "Frankfurt"}, {"Frankfurt", "Stuttgart"},
		{"Berlin", "Rostock"}, {"Berlin", "Magdeburg"},
		{"Hannover", "Bremen"},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	result := g.DFS("Berlin")
	if len(result) == 0 {
		t.Error("DFS returned empty result")
	}

	// Verify first city is Berlin
	if result[0] != "Berlin" {
		t.Errorf("Expected first city to be Berlin, got %s", result[0])
	}
}
