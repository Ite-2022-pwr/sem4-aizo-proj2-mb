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

	err := myGraph.SaveToFile(graph, "graph.txt")
	if err != nil {
		fmt.Println("Error saving graph to file:", err)
	} else {
		fmt.Println("Graph saved to graph.txt")
	}

	readGraph, err := myGraph.ReadFromFile("graph.txt", false, true)
	if err != nil {
		fmt.Println("Error reading graph from file:", err)
	} else {
		fmt.Println("Graph read from graph.txt")
		fmt.Println(readGraph.ToString())
	}

}
