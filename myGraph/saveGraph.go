package myGraph

import (
	"fmt"
	"os"
)

// zapisywanie grafu do pliku
func SaveToFile(graph Graph, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	vertices := graph.GetVertexCount()
	edges := graph.GetEdgeCount()
	_, err = fmt.Fprintf(file, "%d %d\n", vertices, edges)
	if err != nil {
		return err
	}

	allEdges := graph.GetAllEdges()
	for _, edge := range allEdges {
		_, err = fmt.Fprintf(file, "%d %d %d\n", edge.Start, edge.End, edge.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}
