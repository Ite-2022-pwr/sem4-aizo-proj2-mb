package myGraph

import (
	"fmt"
	"log"
	"math"
	"projekt2/timeTrack"
	"projekt2/utils"
	"time"
)

// algorytm Dijkstry do znajdowania najkrótszej ścieżki w grafie
func Dijkstra(inputGraph Graph, startVertex, endVertex int) (path *Path, elapsed int64) {
	name := fmt.Sprintf("Dijkstra, na grafie o %d wierzchołkach, reprezentacja: %s", inputGraph.GetVertexCount(), inputGraph.GetRepresentationName())
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, name)
	}()

	// inicializacja listy dla wyszukiwania najkrótszej ścieżki
	vertexCount := inputGraph.GetVertexCount()
	predecessorDistanceToStartList := make([]VertexPathfinding, vertexCount)

	// ustawienie dla każdego wierzchołka początkowych wartości
	for i := 0; i < vertexCount; i++ {
		predecessorDistanceToStartList[i] = VertexPathfinding{
			Index:         i,
			Predecessor:   -1,
			WeightToStart: math.MaxInt - 1000,
			Visited:       false,
		}
	}

	// ustawienie dla wierzchołka startowego wartości 0
	predecessorDistanceToStartList[startVertex].WeightToStart = 0

	// lista odwiedzonych wierzchołków
	visited := make([]int, 0)

	// rozpoczęcie od wierzchołka startowego
	visitingNow := startVertex
	visitingNowPointer := &predecessorDistanceToStartList[visitingNow]

	// pętla główna
	for len(visited) < vertexCount {
		// wybierz krawędzie wychodzące z wierzchołka visitingNow
		availableEdges := inputGraph.GetAllEdgesFrom(visitingNow)

		// relaksacja krawędzi
		for _, edge := range availableEdges {
			newWeight := visitingNowPointer.WeightToStart + edge.Weight
			checkingVertex := &predecessorDistanceToStartList[edge.End]
			if newWeight < checkingVertex.WeightToStart {
				checkingVertex.WeightToStart = newWeight
				checkingVertex.Predecessor = visitingNow
			}
		}

		// oznacz wierzchołek jako odwiedzony
		visited = append(visited, visitingNow)
		visitingNowPointer.Visited = true

		// wybierz kolejny wierzchołek do odwiedzenia
		minWeight := math.MaxInt
		for i := 0; i < vertexCount; i++ {
			if !predecessorDistanceToStartList[i].Visited && predecessorDistanceToStartList[i].WeightToStart < minWeight && !utils.InListInt(visited, i) {
				minWeight = predecessorDistanceToStartList[i].WeightToStart
				visitingNow = i
				visitingNowPointer = &predecessorDistanceToStartList[i]
			}
		}

		// jeżeli nie ma wierzchołka do odwiedzenia, przerwij pętlę
		if minWeight == math.MaxInt {
			break
		}
	}

	// utworzenie ścieżki
	path = NewPath()
	currentVertex := endVertex
	for currentVertex != startVertex {
		predecessor := predecessorDistanceToStartList[currentVertex].Predecessor
		if predecessor == -1 {
			// jeżeli nie ma poprzednika dla wierzchołka, to nie ma ścieżki
			fmt.Println("No path found from", startVertex, "to", endVertex)
			log.Println("No path found from", startVertex, "to", endVertex)
			return nil, timeTrack.TimeTrack(startTime, "Dijkstra")
		}
		edgeWeight := inputGraph.GetEdgeWeight(predecessor, currentVertex)
		path.AddEdge(Edge{Start: predecessor, End: currentVertex, Weight: edgeWeight})
		currentVertex = predecessor
	}

	// odwracanie krawędzi aby uzyskać ścieżkę od start do end
	for i, j := 0, len(path.Edges)-1; i < j; i, j = i+1, j-1 {
		path.Edges[i], path.Edges[j] = path.Edges[j], path.Edges[i]
	}

	// wypisanie ścieżki
	pathString := path.ToString()
	fmt.Println(pathString)
	log.Println(pathString)

	// zwrócenie ścieżki
	return path, 0
}
