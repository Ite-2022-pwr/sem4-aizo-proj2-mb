package graphRepresentations

type Graph interface {
	IsDirected() bool
	SetDirected()
	GetVertexCount() int
	GetEdgeCount() int
	IsAdjacent() bool
	GetNeighbours() []int
	AddVertex() int
	RemoveVertex()
	GetEdge() int
	AddEdge() int
	RemoveEdge()
	ClearBrokenEdges()
	GetEdgeWeight() int
	SetEdgeWeight()
}
