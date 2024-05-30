package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GeneratePredecessorListGraph(10, 10, true)
	fmt.Println(graph.ToString())

	test, _ := myGraph.Dijkstra(graph, 0)
	fmt.Println(test)
}
