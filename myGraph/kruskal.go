package myGraph

import (
	"fmt"
	"projekt2/timeTrack"
	"projekt2/utils"
	"time"
)

// algorytm kruskala do znajdowania minimalnego drzewa rozpinającego
func Kruskal(inputGraph Graph, outputIncidenceOrPredecessor bool) (mst Graph, elapsed int64) {
	name := fmt.Sprintf("Kruskal, na grafie o %d wierzchołkach, reprezentacja: %s", inputGraph.GetVertexCount(), inputGraph.GetRepresentationName())
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, name)
	}()
	// sortowanie krawędzi niemalejąco
	sortedEdges := SortEdgesListQS(inputGraph.GetAllEdges())

	// inicializacja grafu wynikowego
	if outputIncidenceOrPredecessor {
		mst = NewIncidenceMatrix()
	} else {
		mst = NewPredecessorList()
	}

	// dodanie wierzchołków do grafu wynikowego
	for i := 0; i < inputGraph.GetVertexCount(); i++ {
		mst.AddVertex()
	}

	// utworzenie mapy wierzchołków
	verticesMap := make(map[int]int)
	for i := 0; i < mst.GetVertexCount(); i++ {
		verticesMap[i] = i
	}

	// lista wierzchołków połączonych
	connected := make([]int, 0)

	// iteracja po krawędziach i dodawanie ich do drzewa jeśli nie tworzą cyklu
	for _, edge := range sortedEdges {
		// jeżeli wszystkie wierzchołki są już połączone to kończymy
		if len(connected) == mst.GetVertexCount() {
			break
		}
		start := edge.Start
		end := edge.End

		// sprawdzenie czy krawędź nie tworzy cyklu
		if start != end && !mst.IsAdjacent(start, end) && verticesMap[start] != verticesMap[end] {
			// dodanie krawędzi do drzewa
			mst.AddEdge(start, end, edge.Weight)

			// aktualizacja mapy wierzchołków
			oldComponent := verticesMap[start]
			newComponent := verticesMap[end]
			for vertex, component := range verticesMap {
				if component == oldComponent {
					verticesMap[vertex] = newComponent
				}
			}

			// dodanie wierzchołków do listy połączonych
			if !utils.InListInt(connected, end) {
				connected = append(connected, end)
			}
			if !utils.InListInt(connected, start) {
				connected = append(connected, start)
			}
		}
	}

	// sprawdzenie czy udało się połączyć wszystkie wierzchołki
	if len(connected) != mst.GetVertexCount() && mst.GetEdgeCount() != mst.GetVertexCount()-1 {
		return nil, 0
	}

	return mst, 0
}
