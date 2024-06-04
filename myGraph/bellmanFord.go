package myGraph

import (
	"fmt"
	"log"
	"math"
	"projekt2/timeTrack"
	"time"
)

func BellmanFord(graph Graph, start, end int) (path *Path, elapsed int64) {
	name := fmt.Sprintf("Bellman-Ford, na grafie o %d wierzchołkach, reprezentacja: %s", graph.GetVertexCount(), graph.GetRepresentationName())
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, name)
	}()

	vertexCount := graph.GetVertexCount()
	verticesWithPredecessorsAndWeightToStart := make([]VertexPathfinding, vertexCount)

	// inicjalizacja listy dla wyszukiwania najkrótszej ścieżki
	for i := 0; i < vertexCount; i++ {
		verticesWithPredecessorsAndWeightToStart[i] = VertexPathfinding{
			Index:         i,
			Predecessor:   -1,
			WeightToStart: math.MaxInt32,
			Visited:       false,
		}
	}
	verticesWithPredecessorsAndWeightToStart[start].WeightToStart = 0

	// relaksacja krawędzi
	edges := graph.GetAllEdges()
	for i := 0; i < vertexCount-1; i++ {
		noImprovements := true
		for _, edge := range edges {
			newWeight := verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart + edge.Weight
			if verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart != math.MaxInt32 && newWeight < verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart {
				verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart = newWeight
				verticesWithPredecessorsAndWeightToStart[edge.End].Predecessor = edge.Start
				noImprovements = false
			}
		}
		if noImprovements {
			log.Println("No improvements made on iteration", i+1)
			fmt.Println("No improvements made on iteration", i+1)
			break
		}
	}

	// sprawdzenie czy graf zawiera ujemny cykl
	for _, edge := range edges {
		if verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart != math.MaxInt32 && verticesWithPredecessorsAndWeightToStart[edge.Start].WeightToStart+edge.Weight < verticesWithPredecessorsAndWeightToStart[edge.End].WeightToStart {
			fmt.Println("Graph contains a negative-weight cycle")
			log.Println("Graph contains a negative-weight cycle")
			return nil, 0
		}
	}

	// tworzenie ścieżki
	path = NewPath()
	currentVertex := end
	for currentVertex != start {
		predecessor := verticesWithPredecessorsAndWeightToStart[currentVertex].Predecessor
		if predecessor == -1 {
			// jeżei nie ma ścieżki z start do end
			fmt.Println("No path found from", start, "to", end)
			log.Println("No path found from", start, "to", end)
			return nil, timeTrack.TimeTrack(startTime, "BellmanFord")
		}
		edgeWeight := graph.GetEdgeWeight(predecessor, currentVertex)
		fmt.Println("Adding edge to path:", predecessor, "->", currentVertex, "with weight", edgeWeight)
		log.Println("Adding edge to path:", predecessor, "->", currentVertex, "with weight", edgeWeight)
		path.AddEdge(Edge{Start: predecessor, End: currentVertex, Weight: edgeWeight})
		currentVertex = predecessor
	}

	// odwrócenie krawędzi aby uzyskać ścieżkę od start do end
	for i, j := 0, len(path.Edges)-1; i < j; i, j = i+1, j-1 {
		path.Edges[i], path.Edges[j] = path.Edges[j], path.Edges[i]
	}

	// wyświetlenie ścieżki
	pathString := path.ToString()
	fmt.Println(pathString)
	log.Println(pathString)

	// zwrócenie ścieżki i czasu wykonania
	return path, 0
}
