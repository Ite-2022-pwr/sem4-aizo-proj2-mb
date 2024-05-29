package myGraph

import (
	"math/rand"
	"projekt2/utils"
)

func GenerateGraphIncidenceMatrix(vertices, percentageConnected int, directed bool) *IncidenceMatrix {
	graph := NewIncidenceMatrix()
	graph.SetDirected(directed)
	connected := make([]int, 0)
	for i := 0; i < vertices; i++ {
		graph.AddVertex()
	}
	maxEdges := 0
	if directed {
		maxEdges = vertices * (vertices - 1)
	} else {
		maxEdges = vertices * (vertices - 1) / 2
	}
	edges := maxEdges * percentageConnected / 100
	graph.AddEdge(0, 1, rand.Intn(100)+1)
	connected = append(connected, 0, 1)
	for {
		if len(connected) == vertices {
			break
		}
		start := connected[rand.Intn(len(connected))]
		end := rand.Intn(vertices)
		if utils.InListInt(connected, end) {
			continue
		}
		if start != end && !graph.IsAdjacent(start, end) {
			graph.AddEdge(start, end, rand.Intn(10)+1)
			connected = append(connected, end)
		}
		if graph.GetEdgeCount() == edges {
			break
		}
	}
	for graph.GetEdgeCount() < edges {
		start := rand.Intn(vertices)
		end := rand.Intn(vertices)
		if start != end && !graph.IsAdjacent(start, end) {
			graph.AddEdge(start, end, rand.Intn(10)+1)
		}

	}
	return graph
}

func GeneratePredecessorListGraph(vertices, percentageConnected int, directed bool) *PredecessorList {
	graph := NewPredecessorList()
	graph.SetDirected(directed)
	connected := make([]int, 0)
	for i := 0; i < vertices; i++ {
		graph.AddVertex()
	}
	maxEdges := vertices * (vertices - 1) / 2
	edges := maxEdges * percentageConnected / 100
	graph.AddEdge(0, 1, rand.Intn(10)+1)
	connected = append(connected, 0, 1)
	for {
		if len(connected) == vertices {
			break
		}
		start := connected[rand.Intn(len(connected))]
		end := rand.Intn(vertices)
		if utils.InListInt(connected, end) {
			continue
		}
		if start != end && !graph.IsAdjacent(start, end) {
			graph.AddEdge(start, end, rand.Intn(10)+1)
			connected = append(connected, end)
		}
		if graph.GetEdgeCount() == edges {
			break
		}
	}
	for graph.GetEdgeCount() < edges {
		start := rand.Intn(vertices)
		end := rand.Intn(vertices)
		if start != end && !graph.IsAdjacent(start, end) {
			graph.AddEdge(start, end, rand.Intn(10)+1)
		}

	}
	return graph
}
