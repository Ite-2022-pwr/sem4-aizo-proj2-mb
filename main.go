package main

import (
	"fmt"
	"log"
	"os"
	"projekt2/testSuite"
	"time"
)

func main() {

	//save all logs to file
	dateString := time.Now().Format("2006-01-02_15:04:05")
	logFileName := dateString + ".log"
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error opening file:", err)
	} else {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Println("Error closing file:", err)
			}
		}(f)
		log.SetOutput(f)
	}

	testSuite.RunFullTests()

	//vertices := 100
	//percentageConnected := 25
	//
	//graphDirected := myGraph.GenerateRandomGraph(vertices, percentageConnected, true, true)
	//
	//fmt.Println(graphDirected.GetEdgeCount())
	//
	//myGraph.Dijkstra(graphDirected, 0, 69)
	////test, _ := myGraph.Dijkstra(graph, 0)
	////fmt.Println(test)
	//
	//myGraph.BellmanFord(graphDirected, 0, 69)
	////test2, _ := myGraph.BellmanFord(graph, 0)
	////fmt.Println(test2)
	//
	//err1 := myGraph.SaveToFile(graphDirected, "graphDirectedThousandVertices.txt")
	//if err1 != nil {
	//	fmt.Println("Error saving graph to file:", err1)
	//} else {
	//	fmt.Println("Graph saved")
	//	log.Println("Graph saved")
	//}
	//
	//graphUndirected := myGraph.GenerateRandomGraph(vertices, percentageConnected, false, true)
	//
	//fmt.Println(graphUndirected.GetEdgeCount())
	//
	//myGraph.Kruskal(graphUndirected, true)
	//
	//myGraph.Prim(graphUndirected, 0, true)
	//
	////mst1, time1 := myGraph.Kruskal(graph, true)
	////fmt.Println(mst1.GetEdgeCount())
	////fmt.Println("Time:", time1)
	//////
	////mst2, time1 := myGraph.Prim(graph, 0, true)
	////fmt.Println(mst2.GetEdgeCount())
	////fmt.Println("Time:", time1)
	//
	//err2 := myGraph.SaveToFile(graphUndirected, "graphDirectedThousandVertices.txt")
	//if err2 != nil {
	//	fmt.Println("Error saving graph to file:", err2)
	//} else {
	//	fmt.Println("Graph saved")
	//	log.Println("Graph saved")
	//}

}
