package testSuite

import "projekt2/myGraph"

func SingleTestKruskal(vertices, percentageConnected int) (avgIMTime, avgPLTime float64) {
	// Kruskal's Algorithm
	// Create a random graph
	im := myGraph.GenerateRandomGraph(vertices, percentageConnected, false, true)
	pl := myGraph.NewPredecessorList()
	pl.SetDirected(false)
	for i := 0; i < im.GetVertexCount(); i++ {
		pl.AddVertex()
	}
	for _, edge := range im.GetAllEdges() {
		pl.AddEdge(edge.Start, edge.End, edge.Weight)
	}

	avgIMTime = 0
	avgPLTime = 0
	for i := 0; i < 10; i++ {
		_, timeIM := myGraph.Kruskal(im, true)
		_, timePL := myGraph.Kruskal(pl, false)
		avgIMTime += float64(timeIM)
		avgPLTime += float64(timePL)
	}
	avgIMTime /= 10
	avgPLTime /= 10
	return avgIMTime, avgPLTime
}

func PercentageTestKruskal(percentageConnected int) (avgIMTimes, avgPLTimes map[int]float64) {
	avgIMTimes = make(map[int]float64)
	avgPLTimes = make(map[int]float64)
	for i := 0; i < 7; i++ {
		vertices := 10 + i*10
		avgIMTime, avgPLTime := SingleTestKruskal(vertices, percentageConnected)
		avgIMTimes[vertices] = avgIMTime
		avgPLTimes[vertices] = avgPLTime
	}
	return avgIMTimes, avgPLTimes
}
