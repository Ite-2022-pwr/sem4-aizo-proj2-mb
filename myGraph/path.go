package myGraph

import (
	"fmt"
	"strings"
)

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

func (p *Path) ToString() string {
	if len(p.Edges) == 0 {
		return "Empty path"
	}

	startVertex := p.Edges[0].Start
	endVertex := p.Edges[len(p.Edges)-1].End

	var pathBuilder strings.Builder
	fmt.Fprintf(&pathBuilder, "Path from %d to %d, total weight: %d: ", startVertex, endVertex, p.Weight)
	for i, edge := range p.Edges {
		if i > 0 {
			pathBuilder.WriteString(" -> ")
		}
		fmt.Fprintf(&pathBuilder, "%d-%d->%d", edge.Start, edge.Weight, edge.End)
	}
	return pathBuilder.String()
}
