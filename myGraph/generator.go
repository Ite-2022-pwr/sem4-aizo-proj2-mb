package myGraph

import (
	"fmt"
	"log"
	"math/rand"
	"projekt2/timeTrack"
	"time"
)

// wygeneruj graf losowy
func GenerateRandomGraph(vertices, percentageConnected int, directed bool) (im, pl Graph) {
	defer timeTrack.TimeTrack(time.Now(), "GenerateRandomGraph")

	im = NewIncidenceMatrix()
	pl = NewPredecessorList()

	im.SetDirected(directed)
	pl.SetDirected(directed)
	// dodanie krawędzi
	for i := 0; i < vertices; i++ {
		im.AddVertex()
		pl.AddVertex()
	}

	// obliczenie maksymalnej ilości krawędzi
	maxEdges := vertices * (vertices - 1)
	if !directed {
		maxEdges /= 2
	}

	// oblizenie ilości krawędzi do dodania
	desiredEdges := maxEdges * percentageConnected / 100
	minEdges := vertices - 1
	if desiredEdges < minEdges {
		desiredEdges = minEdges
	}

	// upewnienie się, że graf jest spójny
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

	// dodanie pozostałych krawędzi
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
