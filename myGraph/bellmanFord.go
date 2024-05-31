package myGraph

import (
	"fmt"
	"math"
)

func BellmanFord(graph Graph, start int) (verticesWithPredecessorsAndWeightToStart []VertexPathfinding) {
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
			return nil
		}
	}

	return verticesWithPredecessorsAndWeightToStart
}
