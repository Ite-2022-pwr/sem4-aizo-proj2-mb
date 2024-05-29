package myGraph

import (
	"math"
)

func Dijkstra(inputGraph Graph, startVertex int) (verticesWithPredecessorsAndWeightToStart []VertexPathfinding, err error) {
	predecessorDistanceToStartList := make([]VertexPathfinding, 0)
	vertexCount := inputGraph.GetVertexCount()
	visited := make([]int, 0)
	nextVertices := make([]int, 0)
	nextVertices = append(nextVertices, 0)
	visitingNow := startVertex
	visitingNowPointer := FindByIndex(predecessorDistanceToStartList, visitingNow)
	availableEdges := make([]Edge, 0)
	for i := 0; i < vertexCount; i++ {
		predecessorDistanceToStartList = append(predecessorDistanceToStartList, VertexPathfinding{i, -1, (math.MaxInt - 1000), false})
	}
	predecessorDistanceToStartList[startVertex].WeightToStart = 0
	predecessorDistanceToStartList[startVertex].Visited = true
	for len(visited) < vertexCount {
		availableEdges = inputGraph.GetAllEdgesFrom(visitingNow)
		for _, edge := range availableEdges {
			temp := predecessorDistanceToStartList[edge.Start].WeightToStart + edge.Weight
			checkingVertex := FindByIndex(predecessorDistanceToStartList, edge.End)
			if temp < checkingVertex.WeightToStart { // indexing changes and thats why it fucking breaksS!!!!!!!
				checkingVertex.WeightToStart = temp
				checkingVertex.Predecessor = visitingNow
			}
		}
		visited = append(visited, visitingNow)
		visitingNowPointer.Visited = true
		SortByWeightToStartQS(predecessorDistanceToStartList)
		for j := 0; j < len(predecessorDistanceToStartList); j++ {
			if !predecessorDistanceToStartList[j].Visited {
				visitingNow = predecessorDistanceToStartList[j].Index
				visitingNowPointer = FindByIndex(predecessorDistanceToStartList, visitingNow)
				break
			}
		}
	}
	return predecessorDistanceToStartList, nil
}
