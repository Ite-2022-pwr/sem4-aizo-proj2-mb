package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GenerateGraphIncidenceMatrix(5, 100, false)
	fmt.Println(graph.ToString())

	mst, _ := myGraph.Kruskal(graph, true)
	fmt.Println(mst.ToString())

}
