package myGraph

import (
	"fmt"
	"log"
	"math"
	"projekt2/timeTrack"
	"projekt2/utils"
	"time"
)

// Dijkstra's algorithm to find the shortest path from a start vertex to an end vertex in the graph
func Dijkstra(inputGraph Graph, startVertex, endVertex int) (path *Path, elapsed int64) {
	name := fmt.Sprintf("Dijkstra, na grafie o %d wierzchołkach, reprezentacja: %s", inputGraph.GetVertexCount(), inputGraph.GetRepresentationName())
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, name)
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

	// Step 3: Set the start vertex distance to 0
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
			if !predecessorDistanceToStartList[i].Visited && predecessorDistanceToStartList[i].WeightToStart < minWeight && !utils.InListInt(visited, i) {
				minWeight = predecessorDistanceToStartList[i].WeightToStart
				visitingNow = i
				visitingNowPointer = &predecessorDistanceToStartList[i]
			}
		}

		// If there are no more vertices to visit, break the loop
		if minWeight == math.MaxInt {
			break
		}
	}

	// Construct the path from startVertex to endVertex
	path = NewPath()
	currentVertex := endVertex
	for currentVertex != startVertex {
		predecessor := predecessorDistanceToStartList[currentVertex].Predecessor
		if predecessor == -1 {
			// If no predecessor is found, it means there is no path
			fmt.Println("No path found from", startVertex, "to", endVertex)
			log.Println("No path found from", startVertex, "to", endVertex)
			return nil, timeTrack.TimeTrack(startTime, "Dijkstra")
		}
		edgeWeight := inputGraph.GetEdgeWeight(predecessor, currentVertex)
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
