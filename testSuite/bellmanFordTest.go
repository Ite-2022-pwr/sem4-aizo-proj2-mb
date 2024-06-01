package testSuite

import (
	"fmt"
	"log"
	"math/rand"
	"projekt2/myGraph"
	"runtime/debug"
)

func SingleTestBellmanFord(vertices, percentageConnected int) (avgIMTime, avgPLTime float64) {
	// Bellman-Ford Algorithm
	// Create a random graph
	im, pl := myGraph.GenerateRandomGraph(vertices, percentageConnected, true)

	startVertex := 0

	avgIMTime = 0
	avgPLTime = 0
	for i := 0; i < 10; i++ {
		randomEndVertex := rand.Intn(vertices - 1)
		fmt.Println("Macierz incydencji:")
		log.Println("Macierz incydencji:")
		_, timeIM := myGraph.BellmanFord(im, startVertex, randomEndVertex)
		fmt.Println("--------------------------------")
		log.Println("--------------------------------")
		fmt.Println("Lista poprzedników:")
		log.Println("Lista poprzedników:")
		_, timePL := myGraph.BellmanFord(pl, startVertex, randomEndVertex)
		fmt.Println("--------------------------------")
		log.Println("--------------------------------")
		avgIMTime += float64(timeIM)
		avgPLTime += float64(timePL)
		debug.FreeOSMemory()
	}
	avgIMTime /= 10
	avgPLTime /= 10
	return avgIMTime, avgPLTime
}

func PercentageTestBellmanFord(percentageConnected int) (avgIMTimes, avgPLTimes map[int]float64) {
	avgIMTimes = make(map[int]float64)
	avgPLTimes = make(map[int]float64)
	for i := 0; i < 7; i++ {
		vertices := 100 + i*100
		avgIMTime, avgPLTime := SingleTestBellmanFord(vertices, percentageConnected)
		avgIMTimes[vertices] = avgIMTime
		avgPLTimes[vertices] = avgPLTime
		debug.FreeOSMemory()
	}
	return avgIMTimes, avgPLTimes
}
