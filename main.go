package main

import (
	"fmt"
	"projekt2/myGraph"
)

func main() {
	graphDirected := myGraph.GenerateRandomGraph(1000, 100, true, true)

	fmt.Println(graphDirected.GetEdgeCount())

	myGraph.Dijkstra(graphDirected, 0)
	//test, _ := myGraph.Dijkstra(graph, 0)
	//fmt.Println(test)

	myGraph.BellmanFord(graphDirected, 0)
	//test2, _ := myGraph.BellmanFord(graph, 0)
	//fmt.Println(test2)

	err1 := myGraph.SaveToFile(graphDirected, "graphDirectedThousandVertices.txt")
	if err1 != nil {
		fmt.Println("Error saving graph to file:", err1)
	} else {
		fmt.Println("Graph saved to graph.txt")
	}

	graphUndirected := myGraph.GenerateRandomGraph(1000, 100, false, true)

	fmt.Println(graphUndirected.GetEdgeCount())

	myGraph.Kruskal(graphUndirected, true)

	myGraph.Prim(graphUndirected, 0, true)

	//mst1, time1 := myGraph.Kruskal(graph, true)
	//fmt.Println(mst1.GetEdgeCount())
	//fmt.Println("Time:", time1)
	////
	//mst2, time1 := myGraph.Prim(graph, 0, true)
	//fmt.Println(mst2.GetEdgeCount())
	//fmt.Println("Time:", time1)

	err2 := myGraph.SaveToFile(graphUndirected, "graphDirectedThousandVertices.txt")
	if err2 != nil {
		fmt.Println("Error saving graph to file:", err2)
	} else {
		fmt.Println("Graph saved to graph.txt")
	}

}
