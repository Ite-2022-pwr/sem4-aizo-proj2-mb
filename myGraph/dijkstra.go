package myGraph

import (
	"fmt"
	"math"
	"projekt2/timeTrack"
	"time"
)

// Dijkstra's algorithm to find the shortest paths from a start vertex to all other vertices in the graph
func Dijkstra(inputGraph Graph, startVertex int) (verticesWithPredecessorsAndWeightToStart []VertexPathfinding, elapsed int64) {
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, "Dijkstra")
	}()
	// Step 1: Initialize the list to hold the pathfinding information for each vertex
	vertexCount := inputGraph.GetVertexCount()
	predecessorDistanceToStartList := make([]VertexPathfinding, vertexCount)

	// Step 2: Initialize each vertex with infinite distance and no predecessor
	for i := 0; i < vertexCount; i++ {
		predecessorDistanceToStartList[i] = VertexPathfinding{
			Index:         i,
			Predecessor:   -1,
			WeightToStart: math.MaxInt - 1000,
			Visited:       false,
		}
	}

	// Step 3: Set the start vertex distance to 0 and mark it as visited
	predecessorDistanceToStartList[startVertex].WeightToStart = 0

	// List to keep track of visited vertices
	visited := make([]int, 0)

	// Start with the start vertex
	visitingNow := startVertex
	visitingNowPointer := &predecessorDistanceToStartList[visitingNow]

	// Step 4: Main loop to visit all vertices
	for len(visited) < vertexCount {
		// Get all edges from the currently visiting vertex
		availableEdges := inputGraph.GetAllEdgesFrom(visitingNow)

		// Step 5: Relaxation step - update the shortest path estimates
		for _, edge := range availableEdges {
			newWeight := visitingNowPointer.WeightToStart + edge.Weight
			checkingVertex := &predecessorDistanceToStartList[edge.End]
			if newWeight < checkingVertex.WeightToStart {
				checkingVertex.WeightToStart = newWeight
				checkingVertex.Predecessor = visitingNow
			}
		}

		// Mark the current vertex as visited
		visited = append(visited, visitingNow)
		visitingNowPointer.Visited = true

		// Step 6: Find the next vertex to visit (the one with the smallest distance that hasn't been visited)
		minWeight := math.MaxInt
		for i := 0; i < vertexCount; i++ {
			if !predecessorDistanceToStartList[i].Visited && predecessorDistanceToStartList[i].WeightToStart < minWeight {
				minWeight = predecessorDistanceToStartList[i].WeightToStart
				visitingNow = i
				fmt.Println("Visiting now:", visitingNow)
				visitingNowPointer = &predecessorDistanceToStartList[i]
			}
		}

		// If there are no more vertices to visit, break the loop
		if minWeight == math.MaxInt {
			break
		}
	}

	// Return the list of vertices with their shortest path information
	return predecessorDistanceToStartList, 0
}
