package myGraph

import (
	"math"
	"projekt2/utils"
)

func Dijkstra(inputGraph Graph, startVertex int) (verticesWithPredecessorsAndWeightToStart []VertexPathfinding, err error) {
	predecessorDistanceToStartList := make([]VertexPathfinding, 0)
	vertexCount := inputGraph.GetVertexCount()
	visited := make([]int, 0)
	nextVertices := make([]int, 0)
	nextVertices = append(nextVertices, 0)
	visitingNow := startVertex
	availableEdges := make([]Edge, 0)
	for i := 0; i < vertexCount; i++ {
		predecessorDistanceToStartList = append(predecessorDistanceToStartList, VertexPathfinding{i, -1, math.MaxInt})
	}
	predecessorDistanceToStartList[startVertex].WeightToStart = 0
	availableEdges = append(availableEdges, inputGraph.GetAllEdgesFrom(visitingNow)...)
	for len(visited) < vertexCount {
		SortEdgesListQS(availableEdges)
		for _, edge := range availableEdges {
			if edge.Start != visitingNow {
				break
			} else {
				newWeight := predecessorDistanceToStartList[visitingNow].WeightToStart + edge.Weight
				if newWeight < predecessorDistanceToStartList[edge.End].WeightToStart {
					predecessorDistanceToStartList[edge.End].WeightToStart = newWeight
					predecessorDistanceToStartList[edge.End].Predecessor = visitingNow
				}
				if !utils.InListInt(nextVertices, edge.End) {
					nextVertices = append(nextVertices, edge.End)
					availableEdges = append(availableEdges, inputGraph.GetAllEdgesFrom(edge.End)...)
				}
				availableEdges = RemoveEdgeFromList(edge, availableEdges)
				availableEdges = SortEdgesListQS(availableEdges)
			}

		}
		visited = append(visited, visitingNow)
		visitingNow = availableEdges[0].Start
	}

	return predecessorDistanceToStartList, nil
}
