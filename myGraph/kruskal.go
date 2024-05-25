package myGraph

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
	for i := 0; i < mst.GetVertexCount(); i++ {
		verticesMap[i] = i
	}
	for i := 0; i < len(sortedEdges); i++ {
		start := sortedEdges[i].Start
		end := sortedEdges[i].End
		if start != end && !mst.IsAdjacent(start, end) {
			mst.AddEdge(start, end, sortedEdges[i].Weight)
			for k, v := range verticesMap {
				if v == end {
					verticesMap[k] = start
				}
			}
		}
	}
	return mst, nil
}
