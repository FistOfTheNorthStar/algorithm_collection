package graph

import (
	"fmt"
	"testing"
)

func TestAddVertexB(t *testing.T) {
	g := NewGraphB(15)
	g.AddVertex("Berlin")

	if len(g.GetMapOfVertex()) != 1 {
		t.Errorf("Expected map size of 1, got %d", len(g.GetMapOfVertex()))
	}
}

func TestAddEdgeB(t *testing.T) {
	g := NewGraphB(15)
	city1, city2 := "Berlin", "Leipzig"

	g.AddVertex(city1) // location 0
	g.AddVertex(city2) // location 1
	g.AddEdge(city1, city2)

	adjList := g.GetAdjList()
	if adjList[0][0] != 1 || adjList[1][0] != 0 {
		t.Error("Edge was not properly added")
	}
}

func TestBFS(t *testing.T) {
	g := NewGraphB(15)
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
		{"Hannover", "Bremen"}, {"Stuttgart", "München"},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	result := g.BFS("Berlin")

	if len(result) == 0 {
		t.Error("BFS returned empty result")
	}

	if result[0] != "Berlin" {
		t.Errorf("Expected first city to be Berlin, got %s", result[0])
	}

	// Print BFS traversal for verification
	fmt.Printf("BFS traversal: %v\n", result)
}
