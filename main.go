package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graph := myGraph.GenerateRandomGraph(200, 100, true, true)

	fmt.Println(graph.GetEdgeCount())

	myGraph.Dijkstra(graph, 0)
	//test, _ := myGraph.Dijkstra(graph, 0)
	//fmt.Println(test)

	myGraph.BellmanFord(graph, 0)
	//test2, _ := myGraph.BellmanFord(graph, 0)
	//fmt.Println(test2)

	//mst1, time1 := myGraph.Kruskal(graph, true)
	//fmt.Println(mst1.GetEdgeCount())
	//fmt.Println("Time:", time1)
	////
	//mst2, time1 := myGraph.Prim(graph, 0, true)
	//fmt.Println(mst2.GetEdgeCount())
	//fmt.Println("Time:", time1)

	err := myGraph.SaveToFile(graph, "graph.txt")
	if err != nil {
		fmt.Println("Error saving graph to file:", err)
	} else {
		fmt.Println("Graph saved to graph.txt")
	}

	readGraph, err := myGraph.ReadFromFile("graph.txt", true, true)
	if err != nil {
		fmt.Println("Error reading graph from file:", err)
	} else {
		fmt.Println("Graph read from graph.txt")
		fmt.Println(readGraph.GetEdgeCount())
		//fmt.Println(readGraph.ToString())
	}

}
