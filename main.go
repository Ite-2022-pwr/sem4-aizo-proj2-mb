package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GenerateRandomGraph(10, 50, false, true)
	fmt.Println(graph.ToString())
	fmt.Println(graph.GetEdgeCount())

	//test := myGraph.Dijkstra(graph, 0)
	//fmt.Println(test)
	//
	//test2 := myGraph.BellmanFord(graph, 0)
	//fmt.Println(test2)

	mst1, _ := myGraph.Kruskal(graph, true)
	fmt.Println(mst1.ToString())
	//
	mst2, _ := myGraph.Prim(graph, 0, true)
	fmt.Println(mst2.ToString())
}
