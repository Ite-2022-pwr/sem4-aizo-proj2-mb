package graphRepresentations

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
				if !im.IsDirected() {
					neighbours = append(neighbours, j)
				} else {
					if im.VertexEdgeMatrix[j][i] == -1 {
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

func (im *IncidenceMatrix) GetEdge(start, end int) int {
	for i := 0; i < im.GetEdgeCount(); i++ {
		if im.IsDirected() {
			if im.VertexEdgeMatrix[start][i] == -1 && im.VertexEdgeMatrix[end][i] == 1 {
				return i
			}
		} else {
			if im.VertexEdgeMatrix[start][i] != 0 && im.VertexEdgeMatrix[end][i] != 0 {
				return i
			}
		}
	}
	return -1
}

func (im *IncidenceMatrix) AddEdge(start, end, weight int) int {
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
	return len(im.WeightsList) - 1
}

func (im *IncidenceMatrix) RemoveEdge(start, end int) {
	edge := im.GetEdge(start, end)
	if edge == -1 {
		panic("Edge does not exist")
	}
	tempVertexEdgeMatrix := make([][]int, 0)
	tempWeightsList := make([]int, 0)
	for i := 0; i < im.GetEdgeCount(); i++ {
		if i != edge {
			tempVertexEdgeMatrix = append(tempVertexEdgeMatrix, im.VertexEdgeMatrix[i])
			tempWeightsList = append(tempWeightsList, im.WeightsList[i])
		}
	}
	im.VertexEdgeMatrix = tempVertexEdgeMatrix
	im.WeightsList = tempWeightsList
}

func (im *IncidenceMatrix) ClearBrokenEdges() {
	tempVertexEdgeMatrix := make([][]int, 0)
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
			tempVertexEdgeMatrix = append(tempVertexEdgeMatrix, im.VertexEdgeMatrix[i])
			tempWeightsList = append(tempWeightsList, im.WeightsList[i])
		}
	}
	im.VertexEdgeMatrix = tempVertexEdgeMatrix
	im.WeightsList = tempWeightsList
}

func (im *IncidenceMatrix) GetEdgeWeight(edge int) int {
	return im.WeightsList[edge]
}

func (im *IncidenceMatrix) SetEdgeWeight(edge, weight int) {
	im.WeightsList[edge] = weight
}
