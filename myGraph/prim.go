package myGraph

import "projekt2/utils"

func Prim(inputGraph Graph, startVertex int, incidenceOrPredecessor bool) (mst Graph, err error) {
	if incidenceOrPredecessor {
		mst = NewIncidenceMatrix()
	} else {
		mst = NewPredecessorList()
	}
	verticesCount := inputGraph.GetVertexCount()
	for i := 0; i < verticesCount; i++ {
		mst.AddVertex()
	}
	connected := make([]int, 0)
	connected = append(connected, startVertex)
	usedEdges := make([]Edge, 0)
	for len(connected) < verticesCount {
		edges := make([]Edge, 0)
		for _, connectedVertex := range connected {
			temp := inputGraph.GetAllEdgesFrom(connectedVertex)
			edges = append(edges, temp...)
		}
		edges = SortEdgesListQS(edges)
		for {
			if len(edges) == 0 {
				break
			}
			minEdge := edges[0]
			if !minEdge.isInList(usedEdges, false) && !utils.InListInt(connected, minEdge.End) {
				mst.AddEdge(minEdge.Start, minEdge.End, minEdge.Weight)
				usedEdges = append(usedEdges, minEdge)
				connected = append(connected, minEdge.End)
				break
			} else {
				edges = append(edges[:0], edges[1:]...)
			}
		}
	}
	return mst, nil
}
