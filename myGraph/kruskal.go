package myGraph

import (
	"projekt2/utils"
)

func Kruskal(inputGraph Graph, outputIncidenceOrPredecessor bool) (mst Graph, err error) {
	sortedEdges := SortEdgesListQS(inputGraph.GetAllEdges())
	if outputIncidenceOrPredecessor {
		mst = NewIncidenceMatrix()
	} else {
		mst = NewPredecessorList()
	}
	for i := 0; i < inputGraph.GetVertexCount(); i++ {
		mst.AddVertex()
	}
	verticesMap := make(map[int]int)
	connected := make([]int, 0)
	for i := 0; i < mst.GetVertexCount(); i++ {
		verticesMap[i] = i
	}
	for i := 0; i < len(sortedEdges); i++ {
		if len(connected) == mst.GetVertexCount() {
			break
		}
		start := sortedEdges[i].Start
		end := sortedEdges[i].End
		if start != end && !mst.IsAdjacent(start, end) && verticesMap[start] != verticesMap[end] {
			mst.AddEdge(start, end, sortedEdges[i].Weight)
			if utils.InListInt(connected, end) {
				connected = append(connected, start)
				for vertex, subtree := range verticesMap {
					if subtree == verticesMap[start] {
						verticesMap[vertex] = verticesMap[end]
					}
				}
			} else {
				connected = append(connected, end)
				for vertex, subtree := range verticesMap {
					if subtree == verticesMap[end] {
						verticesMap[vertex] = verticesMap[start]
					}
				}
			}
		}
	}
	return mst, nil
}
