package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GenerateGraphIncidenceMatrix(7, 100, false)
	fmt.Println(graph.ToString())

	mst, _ := myGraph.Prim(graph, 0, true)
	fmt.Println(mst.ToString())

}
