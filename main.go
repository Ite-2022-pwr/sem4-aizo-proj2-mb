package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GeneratePredecessorListGraph(5, 100, true)
	for i := 0; i < len(graph.PredecessorList); i++ {
		for j := 0; j < len(graph.PredecessorList[i]); j++ {
			fmt.Print(graph.PredecessorList[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	graph.RemoveVertex(3)

	for i := 0; i < len(graph.PredecessorList); i++ {
		for j := 0; j < len(graph.PredecessorList[i]); j++ {
			fmt.Print(graph.PredecessorList[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	graph2 := myGraph.GenerateGraphIncidenceMatrix(5, 100, true)
	for i := 0; i < len(graph2.VertexEdgeMatrix); i++ {
		for j := 0; j < len(graph2.VertexEdgeMatrix[i]); j++ {
			fmt.Printf("%3d", graph2.VertexEdgeMatrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(graph2.WeightsList); i++ {
		fmt.Printf("%3d", graph2.WeightsList[i])
	}

	fmt.Println()
	graph2.RemoveVertex(0)
	for i := 0; i < len(graph2.VertexEdgeMatrix); i++ {
		for j := 0; j < len(graph2.VertexEdgeMatrix[i]); j++ {
			fmt.Printf("%3d", graph2.VertexEdgeMatrix[i][j])
		}
		fmt.Println()

	}

	fmt.Println()
	for i := 0; i < len(graph2.WeightsList); i++ {
		fmt.Printf("%3d", graph2.WeightsList[i])
	}
}
