package myGraph

import (
	"fmt"
	"math"
)

func BellmanFord(graph Graph, start int) []VertexPathfinding {
	vertexCount := graph.GetVertexCount()
	pathfinding := make([]VertexPathfinding, vertexCount)

	// Initialize the pathfinding list
	for i := 0; i < vertexCount; i++ {
		pathfinding[i] = VertexPathfinding{
			Index:         i,
			Predecessor:   -1,
			WeightToStart: math.MaxInt32,
			Visited:       false,
		}
	}
	pathfinding[start].WeightToStart = 0

	// Relax edges repeatedly
	for i := 0; i < vertexCount-1; i++ {
		edges := graph.GetAllEdges()
		for _, edge := range edges {
			if pathfinding[edge.Start].WeightToStart != math.MaxInt32 && pathfinding[edge.Start].WeightToStart+edge.Weight < pathfinding[edge.End].WeightToStart {
				pathfinding[edge.End].WeightToStart = pathfinding[edge.Start].WeightToStart + edge.Weight
				pathfinding[edge.End].Predecessor = edge.Start
			}
		}
	}

	// Check for negative-weight cycles
	edges := graph.GetAllEdges()
	for _, edge := range edges {
		if pathfinding[edge.Start].WeightToStart != math.MaxInt32 && pathfinding[edge.Start].WeightToStart+edge.Weight < pathfinding[edge.End].WeightToStart {
			fmt.Println("Graph contains a negative-weight cycle")
			return nil
		}
	}

	return pathfinding
}
