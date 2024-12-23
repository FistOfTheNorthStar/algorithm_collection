package graph

const maxVertex = 15

// Vertex represents a node in the graph
type Vertex struct {
	name    string
	visited bool
}

// Graph represents a graph using adjacency matrix
type Graph struct {
	arrayOfVertex     [maxVertex]*Vertex        // cities
	mapOfVertex       map[string]int            // location mapping
	matrixOfAdjVertex [maxVertex][maxVertex]int // adjacency matrix
	numOfVertices     int
	stack             []int
}

// NewGraph creates a new graph instance
func NewGraph() *Graph {
	g := &Graph{
		mapOfVertex: make(map[string]int),
		stack:       make([]int, 0, maxVertex),
	}

	// Initialize adjacency matrix
	for i := 0; i < maxVertex; i++ {
		for j := 0; j < maxVertex; j++ {
			g.matrixOfAdjVertex[i][j] = 0
		}
	}

	return g
}

// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(name string) {
	vertex := &Vertex{
		name:    name,
		visited: false,
	}
	g.mapOfVertex[name] = g.numOfVertices
	g.arrayOfVertex[g.numOfVertices] = vertex
	g.numOfVertices++
}

// AddEdge adds an edge between two vertices
func (g *Graph) AddEdge(city1, city2 string) {
	start := g.mapOfVertex[city1]
	end := g.mapOfVertex[city2]
	g.matrixOfAdjVertex[start][end] = 1
	g.matrixOfAdjVertex[end][start] = 1
}

// DFS performs depth-first search starting from the given city
func (g *Graph) DFS(city string) []string {
	var result []string
	vertex := g.mapOfVertex[city]
	g.arrayOfVertex[vertex].visited = true
	result = append(result, city)
	g.stack = append(g.stack, vertex)

	for len(g.stack) > 0 {
		adjVertex := g.getAdjVertex(g.stack[len(g.stack)-1])
		if adjVertex != -1 {
			g.arrayOfVertex[adjVertex].visited = true
			result = append(result, g.arrayOfVertex[adjVertex].name)
			g.stack = append(g.stack, adjVertex)
		} else {
			g.stack = g.stack[:len(g.stack)-1]
		}
	}

	// Reset visited flags
	for i := 0; i < g.numOfVertices; i++ {
		g.arrayOfVertex[i].visited = false
	}

	return result
}

// getAdjVertex finds an unvisited adjacent vertex
func (g *Graph) getAdjVertex(vertex int) int {
	for adj := 0; adj < g.numOfVertices; adj++ {
		if g.matrixOfAdjVertex[vertex][adj] == 1 && !g.arrayOfVertex[adj].visited {
			return adj
		}
	}
	return -1
}

// GetMapOfVertex returns the vertex mapping
func (g *Graph) GetMapOfVertex() map[string]int {
	return g.mapOfVertex
}

// GetMatrixOfAdjVertex returns the adjacency matrix
func (g *Graph) GetMatrixOfAdjVertex() [maxVertex][maxVertex]int {
	return g.matrixOfAdjVertex
}
