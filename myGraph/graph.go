package myGraph

type Graph interface {
	IsDirected() bool
	SetDirected()
	GetVertexCount() int
	GetEdgeCount() int
	IsAdjacent() bool
	GetNeighbours() []int
	AddVertex() int
	RemoveVertex()
	GetEdge() Edge
	AddEdge()
	RemoveEdge()
	ClearBrokenEdges()
	GetEdgeWeight() int
	SetEdgeWeight()
}
