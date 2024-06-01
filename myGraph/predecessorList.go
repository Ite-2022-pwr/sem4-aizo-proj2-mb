package myGraph

import "fmt"

type PredecessorList struct {
	PredecessorList [][]Predecessor // map[predecessor]weight
	Directed        bool
}

func NewPredecessorList() *PredecessorList {
	return &PredecessorList{}
}

func (pl *PredecessorList) IsDirected() bool {
	return pl.Directed
}

func (pl *PredecessorList) SetDirected(directed bool) {
	pl.Directed = directed
}

func (pl *PredecessorList) GetVertexCount() int {
	return len(pl.PredecessorList)
}

func (pl *PredecessorList) GetEdgeCount() (count int) {
	count = 0
	temp := make([]Edge, 0)
	if pl.IsDirected() {
		for i := 0; i < pl.GetVertexCount(); i++ {
			count += len(pl.PredecessorList[i])
		}
	} else {
		for i := 0; i < pl.GetVertexCount(); i++ {
			for j := 0; j < len(pl.PredecessorList[i]); j++ {
				tempEdge := Edge{Start: pl.PredecessorList[i][j].Vertex, End: i, Weight: pl.PredecessorList[i][j].Weight}
				if !tempEdge.isInList(temp, pl.IsDirected()) {
					temp = append(temp, tempEdge)
				}
			}
		}
		count = len(temp)
	}
	return count
}

func (pl *PredecessorList) IsAdjacent(vertexA, vertexB int) bool {
	for i := 0; i < len(pl.PredecessorList[vertexB]); i++ {
		if pl.PredecessorList[vertexB][i].Vertex == vertexA {
			return true
		}
	}
	if !pl.IsDirected() {
		for i := 0; i < len(pl.PredecessorList[vertexA]); i++ {
			if pl.PredecessorList[vertexA][i].Vertex == vertexB {
				return true
			}
		}
	}
	return false
}

func (pl *PredecessorList) GetNeighbours(vertex int) (neighbours []int) {
	for i := 0; i < pl.GetVertexCount(); i++ {
		for j := 0; j < len(pl.PredecessorList[i]); j++ {
			if pl.PredecessorList[i][j].Vertex == vertex {
				neighbours = append(neighbours, i)
			}
		}
	}
	return neighbours
}

func (pl *PredecessorList) AddVertex() int {
	pl.PredecessorList = append(pl.PredecessorList, make([]Predecessor, 0))
	return pl.GetVertexCount() - 1
}

func (pl *PredecessorList) RemoveVertex(vertex int) {
	tempList := make([][]Predecessor, 0)
	for i := 0; i < pl.GetVertexCount(); i++ {
		if i != vertex {
			tempPredecessor := make([]Predecessor, 0)
			for j := 0; j < len(pl.PredecessorList[i]); j++ {
				if pl.PredecessorList[i][j].Vertex != vertex {
					if pl.PredecessorList[i][j].Vertex > vertex {
						tempPredecessor = append(tempPredecessor, Predecessor{pl.PredecessorList[i][j].Vertex - 1, pl.PredecessorList[i][j].Weight})
					} else {
						tempPredecessor = append(tempPredecessor, pl.PredecessorList[i][j])
					}
				}
			}
			tempList = append(tempList, tempPredecessor)
		}
	}
	pl.PredecessorList = tempList
}

func (pl *PredecessorList) GetEdge(start, end int) Edge {
	for i := 0; i < len(pl.PredecessorList[end]); i++ {
		if pl.PredecessorList[end][i].Vertex == start {
			return Edge{Start: start, End: end, Weight: pl.PredecessorList[end][i].Weight}
		}
	}
	if !pl.IsDirected() {
		for i := 0; i < len(pl.PredecessorList[start]); i++ {
			if pl.PredecessorList[start][i].Vertex == end {
				return Edge{Start: end, End: start, Weight: pl.PredecessorList[start][i].Weight}
			}
		}
	}
	return Edge{0, 0, 0}
}

func (pl *PredecessorList) GetAllEdges() (edges []Edge) {
	edges = make([]Edge, 0)
	for i := 0; i < pl.GetVertexCount(); i++ {
		for j := 0; j < len(pl.PredecessorList[i]); j++ {
			edge := pl.GetEdge(i, pl.PredecessorList[i][j].Vertex)
			if !edge.isInList(edges, pl.IsDirected()) {
				edges = append(edges, edge)
			}
		}
	}
	return edges
}

func (pl *PredecessorList) GetAllEdgesFrom(vertex int) (edges []Edge) {
	neighbours := pl.GetNeighbours(vertex)
	edges = make([]Edge, 0)
	for _, neighbour := range neighbours {
		edges = append(edges, pl.GetEdge(vertex, neighbour))
	}
	return edges
}

func (pl *PredecessorList) AddEdge(start, end, weight int) {
	if pl.IsAdjacent(start, end) {
		panic("Edge already exists")
	}
	pl.PredecessorList[end] = append(pl.PredecessorList[end], Predecessor{Vertex: start, Weight: weight})
	if !pl.IsDirected() {
		pl.PredecessorList[start] = append(pl.PredecessorList[start], Predecessor{Vertex: end, Weight: weight})
	}
}

func (pl *PredecessorList) RemoveEdge(start, end int) {
	tempList := make([][]Predecessor, 0)
	for i := 0; i < pl.GetVertexCount(); i++ {
		tempPredecessor := make([]Predecessor, 0)
		for j := 0; j < len(pl.PredecessorList[i]); j++ {
			if (pl.PredecessorList[i][j].Vertex != start && j != end && pl.IsDirected()) || (pl.PredecessorList[i][j].Vertex != start && pl.PredecessorList[i][j].Vertex != end && !pl.IsDirected()) {
				tempPredecessor = append(tempPredecessor, pl.PredecessorList[i][j])
			}
		}
		tempList = append(tempList, tempPredecessor)
	}
	pl.PredecessorList = tempList
}

func (pl *PredecessorList) ClearBrokenEdges() {
	return
}

func (pl *PredecessorList) GetEdgeWeight(start, end int) int {
	return pl.GetEdge(start, end).Weight
}

func (pl *PredecessorList) SetEdgeWeight(start, end, weight int) {
	for i := 0; i < len(pl.PredecessorList[end]); i++ {
		if pl.PredecessorList[end][i].Vertex == start {
			pl.PredecessorList[end][i].Weight = weight
		}
	}
	if !pl.IsDirected() {
		for i := 0; i < len(pl.PredecessorList[start]); i++ {
			if pl.PredecessorList[start][i].Vertex == end {
				pl.PredecessorList[start][i].Weight = weight
			}
		}
	}
}

func (pl *PredecessorList) ToString() (out string) {
	for i := 0; i < pl.GetVertexCount(); i++ {
		out += fmt.Sprintln(i, ":", pl.PredecessorList[i])
	}
	return out
}

func (pl *PredecessorList) GetRepresentationName() string {
	return "Predecessor List"
}
