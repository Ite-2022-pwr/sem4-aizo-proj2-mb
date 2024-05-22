package myGraph

type PredecessorList struct {
	predecessorList [][]Predecessor // map[predecessor]weight
	directed        bool
}

func NewPredecessorList() *PredecessorList {
	return &PredecessorList{}
}

func (pl *PredecessorList) IsDirected() bool {
	return pl.directed
}

func (pl *PredecessorList) SetDirected(directed bool) {
	pl.directed = directed
}

func (pl *PredecessorList) GetVertexCount() int {
	return len(pl.predecessorList)
}

func (pl *PredecessorList) GetEdgeCount() (count int) {
	count = 0
	if pl.IsDirected() {
		for i := 0; i < pl.GetVertexCount(); i++ {
			count += len(pl.predecessorList[i])
		}
	} else {
		temp := make([]Edge, 0)
		for i := 0; i < pl.GetVertexCount(); i++ {
			for j := 0; j < len(pl.predecessorList[i]); j++ {
				tempEdge := Edge{start: pl.predecessorList[i][j].Vertex, end: i, weight: pl.predecessorList[i][j].Weight}
				if tempEdge.isInList(temp, pl.IsDirected()) {
					temp = append(temp, tempEdge)
				}
			}
		}
		count = len(temp)
	}
	return count
}

func (pl *PredecessorList) IsAdjacent(vertexA, vertexB int) bool {
	for i := 0; i < len(pl.predecessorList[vertexB]); i++ {
		if pl.predecessorList[vertexB][i].Vertex == vertexA {
			return true
		}
	}
	if !pl.IsDirected() {
		for i := 0; i < len(pl.predecessorList[vertexA]); i++ {
			if pl.predecessorList[vertexA][i].Vertex == vertexB {
				return true
			}
		}
	}
	return false
}

func (pl *PredecessorList) GetNeighbours(vertex int) (neighbours []int) {
	for i := 0; i < pl.GetVertexCount(); i++ {
		for j := 0; j < len(pl.predecessorList[i]); j++ {
			if pl.predecessorList[i][j].Vertex == vertex {
				neighbours = append(neighbours, i)
			}
		}
	}
	return neighbours
}

func (pl *PredecessorList) AddVertex() int {
	pl.predecessorList = append(pl.predecessorList, make([]Predecessor, 0))
	return pl.GetVertexCount() - 1
}

func (pl *PredecessorList) RemoveVertex(vertex int) {
	tempList := make([][]Predecessor, 0)
	for i := 0; i < pl.GetVertexCount(); i++ {
		if i != vertex {
			tempPredecessor := make([]Predecessor, 0)
			for j := 0; j < len(pl.predecessorList[i]); j++ {
				if pl.predecessorList[i][j].Vertex != vertex {
					tempPredecessor = append(tempPredecessor, pl.predecessorList[i][j])
				}
			}
			tempList = append(tempList, tempPredecessor)
		}
	}
	pl.predecessorList = tempList
}

func (pl *PredecessorList) GetEdge(start, end int) Edge {
	for i := 0; i < len(pl.predecessorList[start]); i++ {
		if pl.predecessorList[start][i].Vertex == end {
			return Edge{start: start, end: end, weight: pl.predecessorList[start][i].Weight}
		}
	}
	if !pl.IsDirected() {
		for i := 0; i < len(pl.predecessorList[end]); i++ {
			if pl.predecessorList[end][i].Vertex == start {
				return Edge{start: end, end: start, weight: pl.predecessorList[end][i].Weight}
			}
		}
	}
	return Edge{0, 0, 0}
}

func (pl *PredecessorList) AddEdge(start, end, weight int) {
	if pl.IsAdjacent(start, end) {
		panic("Edge already exists")
	}
	pl.predecessorList[end] = append(pl.predecessorList[end], Predecessor{Vertex: start, Weight: weight})
	if !pl.IsDirected() {
		pl.predecessorList[start] = append(pl.predecessorList[start], Predecessor{Vertex: end, Weight: weight})
	}
}

func (pl *PredecessorList) RemoveEdge(start, end int) {
	tempList := make([][]Predecessor, 0)
	for i := 0; i < pl.GetVertexCount(); i++ {
		tempPredecessor := make([]Predecessor, 0)
		for j := 0; j < len(pl.predecessorList[i]); j++ {
			if (pl.predecessorList[i][j].Vertex != start && j != end && pl.IsDirected()) || (pl.predecessorList[i][j].Vertex != start && pl.predecessorList[i][j].Vertex != end && !pl.IsDirected()) {
				tempPredecessor = append(tempPredecessor, pl.predecessorList[i][j])
			}
		}
		tempList = append(tempList, tempPredecessor)
	}
	pl.predecessorList = tempList
}

func (pl *PredecessorList) ClearBrokenEdges() {
	return
}

func (pl *PredecessorList) GetEdgeWeight(start, end int) int {
	return pl.GetEdge(start, end).weight
}

func (pl *PredecessorList) SetEdgeWeight(start, end, weight int) {
	for i := 0; i < len(pl.predecessorList[end]); i++ {
		if pl.predecessorList[end][i].Vertex == start {
			pl.predecessorList[end][i].Weight = weight
		}
	}
	if !pl.IsDirected() {
		for i := 0; i < len(pl.predecessorList[start]); i++ {
			if pl.predecessorList[start][i].Vertex == end {
				pl.predecessorList[start][i].Weight = weight
			}
		}
	}
}
