package myGraph

type Path struct {
	Edges  []Edge
	Weight int
}

func NewPath() *Path {
	return &Path{}
}

func (p *Path) GetEdges() []Edge {
	return p.Edges
}

func (p *Path) GetWeight() int {
	return p.Weight
}

func (p *Path) AddEdge(edge Edge) {
	p.Edges = append(p.Edges, edge)
	p.Weight += edge.Weight
}
