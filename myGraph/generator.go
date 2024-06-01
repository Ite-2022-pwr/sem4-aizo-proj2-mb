package myGraph

import (
	"fmt"
	"log"
	"math/rand"
	"projekt2/timeTrack"
	"time"
)

// GenerateRandomGraph creates a random graph with the specified vertex count, edge infill percentage, directionality, and graph type.
func GenerateRandomGraph(vertices, percentageConnected int, directed bool) (im, pl Graph) {
	defer timeTrack.TimeTrack(time.Now(), "GenerateRandomGraph")

	im = NewIncidenceMatrix()
	pl = NewPredecessorList()

	im.SetDirected(directed)
	pl.SetDirected(directed)
	// Add vertices to the graph
	for i := 0; i < vertices; i++ {
		im.AddVertex()
		pl.AddVertex()
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
		if start != end && connected[start] && !connected[end] && !im.IsAdjacent(start, end) && !pl.IsAdjacent(start, end) {
			newEdge := Edge{Start: start, End: end, Weight: rand.Intn(1000) + 1}
			im.AddEdge(start, end, newEdge.Weight)
			pl.AddEdge(start, end, newEdge.Weight)
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
		if start != end && !im.IsAdjacent(start, end) && !pl.IsAdjacent(start, end) {
			newEdge := Edge{Start: start, End: end, Weight: rand.Intn(1000) + 1}
			im.AddEdge(start, end, newEdge.Weight)
			pl.AddEdge(start, end, newEdge.Weight)
			edgesAdded++
			fmt.Println("Edges added:", edgesAdded, " / Desired:", desiredEdges)
			log.Println("Edges added:", edgesAdded, " / Desired:", desiredEdges)
		}
	}

	return im, pl
}
