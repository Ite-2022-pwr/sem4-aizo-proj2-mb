package myGraph

import (
	"fmt"
	"projekt2/utils"
)

type IncidenceMatrix struct {
	VertexEdgeMatrix [][]int
	WeightsList      []int
	directed         bool
}

func NewIncidenceMatrix() *IncidenceMatrix {
	return &IncidenceMatrix{}
}

func (im *IncidenceMatrix) IsDirected() bool {
	return im.directed
}

func (im *IncidenceMatrix) SetDirected(directed bool) {
	im.directed = directed
}

func (im *IncidenceMatrix) GetVertexCount() int {
	return len(im.VertexEdgeMatrix)
}

func (im *IncidenceMatrix) GetEdgeCount() int {
	return len(im.WeightsList)
}

func (im *IncidenceMatrix) IsAdjacent(vertexA, vertexB int) bool {
	for i := 0; i < im.GetEdgeCount(); i++ {
		if im.VertexEdgeMatrix[vertexA][i] != 0 && im.VertexEdgeMatrix[vertexB][i] != 0 {
			if !im.IsDirected() {
				return true
			} else {
				return im.VertexEdgeMatrix[vertexA][i] == -1 && im.VertexEdgeMatrix[vertexB][i] == 1
			}
		}
	}
	return false
}

func (im *IncidenceMatrix) GetNeighbours(vertex int) (neighbours []int) {
	for i := 0; i < im.GetEdgeCount(); i++ {
		for j := 0; j < im.GetVertexCount(); j++ {
			if im.VertexEdgeMatrix[j][i] != 0 && j != vertex {
				if !im.IsDirected() && !utils.InListInt(neighbours, j) {
					neighbours = append(neighbours, j)
				} else {
					if im.VertexEdgeMatrix[j][i] == 1 && im.VertexEdgeMatrix[vertex][i] == -1 {
						neighbours = append(neighbours, j)
					}
				}
			}
		}
	}
	return neighbours
}

func (im *IncidenceMatrix) AddVertex() int {
	im.VertexEdgeMatrix = append(im.VertexEdgeMatrix, make([]int, im.GetEdgeCount()))
	return im.GetVertexCount() - 1
}

func (im *IncidenceMatrix) RemoveVertex(vertex int) {
	tempMatrix := make([][]int, 0)
	for i := 0; i < im.GetVertexCount(); i++ {
		if i != vertex {
			tempMatrix = append(tempMatrix, im.VertexEdgeMatrix[i])
		}
	}
	im.VertexEdgeMatrix = tempMatrix
	im.ClearBrokenEdges()
}

func (im *IncidenceMatrix) GetEdge(start, end int) Edge {
	for i := 0; i < im.GetEdgeCount(); i++ {
		if im.IsDirected() {
			if im.VertexEdgeMatrix[start][i] == -1 && im.VertexEdgeMatrix[end][i] == 1 {
				return Edge{Start: start, End: end, Weight: im.WeightsList[i]}
			}
		} else {
			if im.VertexEdgeMatrix[start][i] != 0 && im.VertexEdgeMatrix[end][i] != 0 {
				return Edge{Start: start, End: end, Weight: im.WeightsList[i]}
			}
		}
	}
	return Edge{0, 0, 0}
}

func (im *IncidenceMatrix) GetAllEdges() (edges []Edge) {
	for i := 0; i < im.GetEdgeCount(); i++ {
		start := 0
		end := 0
		for j := 0; j < im.GetVertexCount(); j++ {
			if im.IsDirected() && im.VertexEdgeMatrix[j][i] == -1 {
				start = j
			} else if im.VertexEdgeMatrix[j][i] != 0 {
				if start == 0 {
					start = j
				} else {
					end = j
				}
			}

		}
		edges = append(edges, Edge{Start: start, End: end, Weight: im.WeightsList[i]})
	}
	return edges
}

func (im *IncidenceMatrix) GetAllEdgesFrom(vertex int) (edges []Edge) {
	neighbours := im.GetNeighbours(vertex)
	edges = make([]Edge, 0)
	for _, neighbour := range neighbours {
		edges = append(edges, im.GetEdge(vertex, neighbour))
	}
	return edges
}

func (im *IncidenceMatrix) AddEdge(start, end, weight int) {
	if im.IsAdjacent(start, end) {
		panic("Edge already exists")
	}
	for i := 0; i < im.GetVertexCount(); i++ {
		im.VertexEdgeMatrix[i] = append(im.VertexEdgeMatrix[i], 0)
	}
	if im.directed {
		im.VertexEdgeMatrix[start][im.GetEdgeCount()] = -1
		im.VertexEdgeMatrix[end][im.GetEdgeCount()] = 1
	} else {
		im.VertexEdgeMatrix[start][im.GetEdgeCount()] = 1
		im.VertexEdgeMatrix[end][im.GetEdgeCount()] = 1
	}
	im.WeightsList = append(im.WeightsList, weight)
}

func (im *IncidenceMatrix) RemoveEdge(start, end int) {
	tempVertexEdgeMatrix := make([][]int, 0)
	tempWeightsList := make([]int, 0)
	for i := 0; i < im.GetEdgeCount(); i++ {
		thisStart := false
		thisEnd := false
		for j := 0; j < im.GetVertexCount(); j++ {
			if im.VertexEdgeMatrix[j][i] == -1 && j == start {
				thisStart = true
			} else if im.VertexEdgeMatrix[j][i] == 1 && j == end {
				thisEnd = true
			} else if im.VertexEdgeMatrix[j][i] != 0 && !im.IsDirected() {
				if !thisStart {
					thisStart = true
				} else {
					thisEnd = true
				}
			}
		}
		if !thisStart && !thisEnd {
			tempWeightsList = append(tempWeightsList, im.WeightsList[i])
			tempVertexEdgeMatrix = append(tempVertexEdgeMatrix, im.VertexEdgeMatrix[i])
		}
	}
	im.VertexEdgeMatrix = tempVertexEdgeMatrix
	im.WeightsList = tempWeightsList
}

func (im *IncidenceMatrix) ClearBrokenEdges() {
	tempVertexEdgeMatrix := make([][]int, im.GetVertexCount())
	tempWeightsList := make([]int, 0)
	for i := 0; i < im.GetEdgeCount(); i++ {
		hasStart := false
		hasEnd := false
		for j := 0; j < im.GetVertexCount(); j++ {
			if im.VertexEdgeMatrix[j][i] == -1 && im.IsDirected() {
				hasStart = true
			} else if im.VertexEdgeMatrix[j][i] == 1 && im.IsDirected() {
				hasEnd = true
			} else if im.VertexEdgeMatrix[j][i] != 0 && !im.IsDirected() {
				if !hasStart {
					hasStart = true
				} else {
					hasEnd = true
				}
			}
		}
		if hasStart && hasEnd {
			for j := 0; j < im.GetVertexCount(); j++ {
				tempVertexEdgeMatrix[j] = append(tempVertexEdgeMatrix[j], im.VertexEdgeMatrix[j][i])
			}
			tempWeightsList = append(tempWeightsList, im.WeightsList[i])
		}
	}
	im.VertexEdgeMatrix = tempVertexEdgeMatrix
	im.WeightsList = tempWeightsList
}

func (im *IncidenceMatrix) GetEdgeWeight(start, end int) int {
	return im.GetEdge(start, end).Weight
}

func (im *IncidenceMatrix) SetEdgeWeight(start, end, weight int) {
	for i := 0; i < im.GetEdgeCount(); i++ {
		correctStart := false
		correctEnd := false
		for j := 0; j < im.GetVertexCount(); j++ {
			if im.VertexEdgeMatrix[j][i] == -1 && j == start {
				correctStart = true
			} else if im.VertexEdgeMatrix[j][i] == 1 && j == end {
				correctEnd = true
			} else if im.VertexEdgeMatrix[j][i] != 0 && !im.IsDirected() {
				if !correctStart {
					correctStart = true
				} else {
					correctEnd = true
				}
			}
		}
		if correctStart && correctEnd {
			im.WeightsList[i] = weight
		}
	}
}

func (im *IncidenceMatrix) ToString() (out string) {
	for i := 0; i < im.GetVertexCount(); i++ {
		out += fmt.Sprintf("%3d \t:", i)
		for j := 0; j < im.GetEdgeCount(); j++ {
			out += fmt.Sprintf("%3d", im.VertexEdgeMatrix[i][j])
		}
		out += fmt.Sprintf("\n")
	}
	out += fmt.Sprintf("Weights\t:")
	for i := 0; i < im.GetEdgeCount(); i++ {
		out += fmt.Sprintf("%3d", im.WeightsList[i])
	}
	out += fmt.Sprintf("\n")
	return out
}
