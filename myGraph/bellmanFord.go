package myGraph

import (
	"fmt"
	"log"
	"math"
	"projekt2/timeTrack"
	"time"
)

func BellmanFord(graph Graph, start int) (verticesWithPredecessorsAndWeightToStart []VertexPathfinding, elapsed int64) {
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, "BellmanFord")
	}()
	vertexCount := graph.GetVertexCount()
	verticesWithPredecessorsAndWeightToStart = make([]VertexPathfinding, vertexCount)

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
			fmt.Println("Edge:", edge.Start, edge.End, edge.Weight, verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart, verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart)
			log.Println("Edge:", edge.Start, edge.End, edge.Weight, verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart, verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart)
			if verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart != math.MaxInt32 && verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart+edge.Weight < verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart {
				verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart = verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart + edge.Weight
				verticesWithPredecessorsAndWeightToStart[edge.End].Predecessor = edge.Start
			}
		}
	}

	// Check for negative-weight cycles
	edges := graph.GetAllEdges()
	for _, edge := range edges {
		if verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart != math.MaxInt32 && verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart+edge.Weight < verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart {
			fmt.Println("Graph contains a negative-weight cycle")
			return nil, 0
		}
	}

	return verticesWithPredecessorsAndWeightToStart, 0
}
