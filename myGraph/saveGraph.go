package myGraph

import (
	"fmt"
	"os"
)

// SaveToFile saves the graph to a file in the specified format.
func SaveToFile(graph Graph, filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the number of vertices and edges
	vertices := graph.GetVertexCount()
	edges := graph.GetEdgeCount()
	_, err = fmt.Fprintf(file, "%d %d\n", vertices, edges)
	if err != nil {
		return err
	}

	// Write each edge
	allEdges := graph.GetAllEdges()
	for _, edge := range allEdges {
		_, err = fmt.Fprintf(file, "%d %d %d\n", edge.Start, edge.End, edge.Weight)
		if err != nil {
			return err
		}
	}

	return nil
}
