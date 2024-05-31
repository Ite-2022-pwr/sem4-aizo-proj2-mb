package myGraph

import (
	"fmt"
	"log"
	"math/rand"
	"projekt2/timeTrack"
	"time"
)

// GenerateRandomGraph creates a random graph with the specified vertex count, edge infill percentage, directionality, and graph type.
func GenerateRandomGraph(vertices, percentageConnected int, directed, useIncidenceMatrix bool) Graph {
	defer timeTrack.TimeTrack(time.Now(), "GenerateRandomGraph")
	var graph Graph
	if useIncidenceMatrix {
		graph = NewIncidenceMatrix()
	} else {
		graph = NewPredecessorList()
	}
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

	// Calculate the desired number of edges based on the percentage connected
	desiredEdges := maxEdges * percentageConnected / 100
	minEdges := vertices - 1
	if desiredEdges < minEdges {
		desiredEdges = minEdges
	}

	// Ensure the graph is connected using a MST approach
	connected := make([]bool, vertices)
	connected[0] = true
	edgesAdded := 0

	for edgesAdded < minEdges {
		start := rand.Intn(vertices)
		end := rand.Intn(vertices)
		if start != end && connected[start] && !connected[end] && !graph.IsAdjacent(start, end) {
			graph.AddEdge(start, end, rand.Intn(1000)+1)
			connected[end] = true
			edgesAdded++
			fmt.Println("Edges added:", edgesAdded, " / Desired:", desiredEdges)
			log.Println("Edges added:", edgesAdded, " / Desired:", desiredEdges)
		}
	}

	// Add additional random edges to meet the desired edge count
	for edgesAdded < desiredEdges {
		start := rand.Intn(vertices)
		end := rand.Intn(vertices)
		if start != end && !graph.IsAdjacent(start, end) {
			graph.AddEdge(start, end, rand.Intn(10)+1)
			edgesAdded++
			fmt.Println("Edges added:", edgesAdded, " / Desired:", desiredEdges)
			log.Println("Edges added:", edgesAdded, " / Desired:", desiredEdges)
		}
	}

	return graph
}
