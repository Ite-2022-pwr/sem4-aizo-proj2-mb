package myGraph

type Graph interface {
	IsDirected() bool
	SetDirected(bool)
	GetVertexCount() int
	GetEdgeCount() int
	IsAdjacent(int, int) bool
	GetNeighbours(int) []int
	AddVertex() int
	RemoveVertex(int)
	GetEdge(int, int) Edge
	GetAllEdges() []Edge
	GetAllEdgesFrom(int) []Edge
	AddEdge(int, int, int)
	RemoveEdge(int, int)
	ClearBrokenEdges()
	GetEdgeWeight(int, int) int
	SetEdgeWeight(int, int, int)
	ToString() string
}
