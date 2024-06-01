package myGraph

import (
	"fmt"
	"log"
	"math"
	"projekt2/timeTrack"
	"time"
)

func BellmanFord(graph Graph, start, end int) (path *Path, elapsed int64) {
	name := fmt.Sprintf("Bellman-Ford, na grafie o %d wierzchołkach, reprezentacja: %s", graph.GetVertexCount(), graph.GetRepresentationName())
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, name)
	}()

	vertexCount := graph.GetVertexCount()
	verticesWithPredecessorsAndWeightToStart := make([]VertexPathfinding, vertexCount)

	// Initialize the pathfinding list
	for i := 0; i < vertexCount; i++ {
		verticesWithPredecessorsAndWeightToStart[i] = VertexPathfinding{
			Index:         i,
			Predecessor:   -1,
			WeightToStart: math.MaxInt32,
			Visited:       false,
		}
	}
	verticesWithPredecessorsAndWeightToStart[start].WeightToStart = 0

	// Relax edges repeatedly
	for i := 0; i < vertexCount-1; i++ {
		edges := graph.GetAllEdges()
		for _, edge := range edges {
			newWeight := verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart + edge.Weight
			if verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart != math.MaxInt32 && newWeight < verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart {
				verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart = newWeight
				verticesWithPredecessorsAndWeightToStart[edge.End].Predecessor = edge.Start
			}
		}
	}

	// Check for negative-weight cycles
	edges := graph.GetAllEdges()
	for _, edge := range edges {
		if verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart != math.MaxInt32 && verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart+edge.Weight < verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart {
			fmt.Println("Graph contains a negative-weight cycle")
			log.Println("Graph contains a negative-weight cycle")
			return nil, 0
		}
	}

	// Construct the path from start to end
	path = NewPath()
	currentVertex := end
	for currentVertex != start {
		predecessor := verticesWithPredecessorsAndWeightToStart[currentVertex].Predecessor
		if predecessor == -1 {
			// If no predecessor is found, it means there is no path
			fmt.Println("No path found from", start, "to", end)
			log.Println("No path found from", start, "to", end)
			return nil, timeTrack.TimeTrack(startTime, "BellmanFord")
		}
		edgeWeight := graph.GetEdgeWeight(predecessor, currentVertex)
		fmt.Println("Adding edge to path:", predecessor, "->", currentVertex, "with weight", edgeWeight)
		log.Println("Adding edge to path:", predecessor, "->", currentVertex, "with weight", edgeWeight)
		path.AddEdge(Edge{Start: predecessor, End: currentVertex, Weight: edgeWeight})
		currentVertex = predecessor
	}

	// Reverse the edges to get the path from start to end
	for i, j := 0, len(path.Edges)-1; i < j; i, j = i+1, j-1 {
		path.Edges[i], path.Edges[j] = path.Edges[j], path.Edges[i]
	}

	// Log the final path
	pathString := path.ToString()
	fmt.Println(pathString)
	log.Println(pathString)

	// Return the path and the elapsed time
	return path, 0
}
