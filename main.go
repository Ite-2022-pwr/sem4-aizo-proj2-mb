package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GenerateGraphIncidenceMatrix(20, 10, true)
	for i := 0; i < len(graph.VertexEdgeMatrix); i++ {
		for j := 0; j < graph.GetEdgeCount(); j++ {
			fmt.Printf("%3d", graph.VertexEdgeMatrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(graph.WeightsList); i++ {
		fmt.Printf("%3d", graph.WeightsList[i])
	}
}
