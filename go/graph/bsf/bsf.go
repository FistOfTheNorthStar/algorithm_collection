package graph

// Vertex represents a node in the graph
type VertexB struct {
	name    string
	visited bool
}

// Graph represents a graph using adjacency lists
type GraphB struct {
	maxVertex     int
	arrayOfVertex []*VertexB     // cities
	mapOfVertex   map[string]int // vertex mapping
	adjList       [][]int        // adjacency lists
	numOfVertices int
	queue         []int
}

// NewGraph creates a new graph instance
func NewGraphB(vertices int) *GraphB {
	g := &GraphB{
		maxVertex:     vertices,
		arrayOfVertex: make([]*VertexB, vertices),
		mapOfVertex:   make(map[string]int),
		adjList:       make([][]int, vertices),
		queue:         make([]int, 0),
	}

	// Initialize adjacency lists
	for i := 0; i < vertices; i++ {
		g.adjList[i] = make([]int, 0)
	}

	return g
}

// AddVertex adds a new vertex to the graph
func (g *GraphB) AddVertex(city string) {
	vertex := &VertexB{
		name:    city,
		visited: false,
	}
	g.mapOfVertex[city] = g.numOfVertices
	g.arrayOfVertex[g.numOfVertices] = vertex
	g.numOfVertices++
}

// AddEdge adds an edge between two vertices
func (g *GraphB) AddEdge(city1, city2 string) {
	start := g.mapOfVertex[city1]
	end := g.mapOfVertex[city2]
	g.adjList[start] = append(g.adjList[start], end)
	g.adjList[end] = append(g.adjList[end], start)
}

// BFS performs breadth-first search starting from the given city
func (g *GraphB) BFS(city string) []string {
	var result []string
	vertex := g.mapOfVertex[city]
	g.arrayOfVertex[vertex].visited = true
	result = append(result, city)
	g.queue = append(g.queue, vertex)

	for len(g.queue) > 0 {
		headVertex := g.queue[0]
		g.queue = g.queue[1:]

		// Process all unvisited adjacent vertices
		for {
			adjVertex := g.getAdjVertexB(headVertex)
			if adjVertex == -1 {
				break
			}
			g.arrayOfVertex[adjVertex].visited = true
			result = append(result, g.arrayOfVertex[adjVertex].name)
			g.queue = append(g.queue, adjVertex)
		}
	}

	// Reset visited flags
	for i := 0; i < g.numOfVertices; i++ {
		g.arrayOfVertex[i].visited = false
	}

	return result
}

// getAdjVertex finds an unvisited adjacent vertex
func (g *GraphB) getAdjVertexB(vertex int) int {
	for _, adj := range g.adjList[vertex] {
		if !g.arrayOfVertex[adj].visited {
			return adj
		}
	}
	return -1
}

// GetMapOfVertex returns the vertex mapping
func (g *GraphB) GetMapOfVertex() map[string]int {
	return g.mapOfVertex
}

// GetAdjList returns the adjacency lists
func (g *GraphB) GetAdjList() [][]int {
	return g.adjList
}
