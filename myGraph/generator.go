package myGraph

import (
	"math/rand"
)

// GenerateGraphIncidenceMatrix creates a random graph with the specified vertex count, edge infill percentage, and directionality using an incidence matrix representation.
func GenerateGraphIncidenceMatrix(vertices, percentageConnected int, directed bool) *IncidenceMatrix {
	graph := NewIncidenceMatrix()
	graph.SetDirected(directed)

	// Add vertices to the graph
	for i := 0; i < vertices; i++ {
		graph.AddVertex()
	}

	// Calculate the maximum possible edges based on the graph being directed or not
	maxEdges := vertices * (vertices - 1)
	if !directed {
		maxEdges /= 2
	}

	// Calculate the number of edges based on the percentage connected
	desiredEdges := maxEdges * percentageConnected / 100

	// Create a fully connected graph
	for i := 0; i < vertices; i++ {
		for j := i + 1; j < vertices; j++ {
			if directed {
				graph.AddEdge(i, j, rand.Intn(100)+1)
				graph.AddEdge(j, i, rand.Intn(100)+1)
			} else {
				graph.AddEdge(i, j, rand.Intn(100)+1)
			}
		}
	}

	// Randomly remove edges until the desired percentage is reached
	for graph.GetEdgeCount() > desiredEdges {
		edges := graph.GetAllEdges()
		removeIndex := rand.Intn(len(edges))
		removeEdge := edges[removeIndex]

		// Check if removing this edge will disconnect the graph
		graph.RemoveEdge(removeEdge.Start, removeEdge.End)
		if !IsGraphConnected(graph) {
			graph.AddEdge(removeEdge.Start, removeEdge.End, removeEdge.Weight)
		}
	}

	return graph
}

// GeneratePredecessorListGraph creates a random graph with the specified vertex count, edge infill percentage, and directionality using a predecessor list representation.
func GeneratePredecessorListGraph(vertices, percentageConnected int, directed bool) *PredecessorList {
	graph := NewPredecessorList()
	graph.SetDirected(directed)

	// Add vertices to the graph
	for i := 0; i < vertices; i++ {
		graph.AddVertex()
	}

	// Calculate the maximum possible edges based on the graph being directed or not
	maxEdges := vertices * (vertices - 1)
	if !directed {
		maxEdges /= 2
	}

	// Calculate the number of edges based on the percentage connected
	desiredEdges := maxEdges * percentageConnected / 100

	// Create a fully connected graph
	for i := 0; i < vertices; i++ {
		for j := i + 1; j < vertices; j++ {
			if directed {
				graph.AddEdge(i, j, rand.Intn(100)+1)
				graph.AddEdge(j, i, rand.Intn(100)+1)
			} else {
				graph.AddEdge(i, j, rand.Intn(100)+1)
			}
		}
	}

	// Randomly remove edges until the desired percentage is reached
	for graph.GetEdgeCount() > desiredEdges {
		edges := graph.GetAllEdges()
		removeIndex := rand.Intn(len(edges))
		removeEdge := edges[removeIndex]

		// Check if removing this edge will disconnect the graph
		graph.RemoveEdge(removeEdge.Start, removeEdge.End)
		if !IsGraphConnected(graph) {
			graph.AddEdge(removeEdge.Start, removeEdge.End, removeEdge.Weight)
		}
	}

	return graph
}

// IsGraphConnected checks if the graph is connected
func IsGraphConnected(graph Graph) bool {
	visited := make(map[int]bool)
	vertices := graph.GetVertexCount()

	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		for _, neighbor := range graph.GetNeighbours(v) {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	// Start DFS from the first vertex
	dfs(0)

	return len(visited) == vertices
}
