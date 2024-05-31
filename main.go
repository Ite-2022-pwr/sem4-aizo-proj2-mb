package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GenerateGraphIncidenceMatrix(7, 100, false)
	fmt.Println(graph.ToString())

	test, _ := myGraph.Dijkstra(graph, 0)
	fmt.Println(test)

	mst1, _ := myGraph.Kruskal(graph, true)
	fmt.Println(mst1.ToString())

	mst2, _ := myGraph.Prim(graph, 0, true)
	fmt.Println(mst2.ToString())
}
