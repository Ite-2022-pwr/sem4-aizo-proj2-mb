package testSuite

import (
	"math/rand"
	"projekt2/myGraph"
)

func SingleTestBellmanFord(vertices, percentageConnected int) (avgIMTime, avgPLTime float64) {
	// Bellman-Ford Algorithm
	// Create a random graph
	im := myGraph.GenerateRandomGraph(vertices, percentageConnected, true, true)
	pl := myGraph.NewPredecessorList()
	pl.SetDirected(true)
	for i := 0; i < im.GetVertexCount(); i++ {
		pl.AddVertex()
	}
	for _, edge := range im.GetAllEdges() {
		if pl.IsAdjacent(edge.Start, edge.End) {
			continue
		}
		pl.AddEdge(edge.Start, edge.End, edge.Weight)
	}

	startVertex := 0

	avgIMTime = 0
	avgPLTime = 0
	for i := 0; i < 10; i++ {
		randomEndVertex := rand.Intn(vertices - 1)
		_, timeIM := myGraph.BellmanFord(im, startVertex, randomEndVertex)
		_, timePL := myGraph.BellmanFord(pl, startVertex, randomEndVertex)
		avgIMTime += float64(timeIM)
		avgPLTime += float64(timePL)
	}
	avgIMTime /= 10
	avgPLTime /= 10
	return avgIMTime, avgPLTime
}

func PercentageTestBellmanFord(percentageConnected int) (avgIMTimes, avgPLTimes map[int]float64) {
	avgIMTimes = make(map[int]float64)
	avgPLTimes = make(map[int]float64)
	for i := 0; i < 7; i++ {
		vertices := 10 + i*10
		avgIMTime, avgPLTime := SingleTestBellmanFord(vertices, percentageConnected)
		avgIMTimes[vertices] = avgIMTime
		avgPLTimes[vertices] = avgPLTime
	}
	return avgIMTimes, avgPLTimes
}
