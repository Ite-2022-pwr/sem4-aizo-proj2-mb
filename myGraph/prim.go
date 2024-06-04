package myGraph

import (
	"container/heap"
	"fmt"
	"projekt2/timeTrack"
	"time"
)

// Kolejka priorytetowa dla krawędzi
type EdgePriorityQueue []*Edge

func (pq EdgePriorityQueue) Len() int { return len(pq) }

func (pq EdgePriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq EdgePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *EdgePriorityQueue) Push(x interface{}) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}

func (pq *EdgePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *EdgePriorityQueue) ToString() (out string) {
	out = fmt.Sprintf("Priority Queue:\n")
	for _, edge := range *pq {
		out += fmt.Sprintf("%d -> %d: %d\n", edge.Start, edge.End, edge.Weight)
	}
	return out
}

// algorithm Prim do znajdowania minimalnego drzewa rozpinającego
func Prim(inputGraph Graph, startVertex int, incidenceOrPredecessor bool) (mst Graph, elapsed int64) {
	name := fmt.Sprintf("Prim, na grafie o %d wierzchołkach, reprezentacja: %s", inputGraph.GetVertexCount(), inputGraph.GetRepresentationName())
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, name)
	}()
	// inicializacja grafu wynikowego
	if incidenceOrPredecessor {
		mst = NewIncidenceMatrix()
	} else {
		mst = NewPredecessorList()
	}

	verticesCount := inputGraph.GetVertexCount()
	for i := 0; i < verticesCount; i++ {
		mst.AddVertex()
	}

	// utworzenie kolejki priorytetowej dla krawędzi
	pq := &EdgePriorityQueue{}
	heap.Init(pq)

	// lista do sprawdzania czy wierzchołek jest już w MST
	inMST := make([]bool, verticesCount)
	inMST[startVertex] = true

	// dodać wszystkie krawędzie wychodzące z wierzchołka startowego do kolejki priorytetowej
	for _, edge := range inputGraph.GetAllEdgesFrom(startVertex) {
		heap.Push(pq, &edge)
	}

	// przetwarzanie krawędzi
	for pq.Len() > 0 {
		// wybierz krawędź o najmniejszej wadze
		minEdge := heap.Pop(pq).(*Edge)
		if !inMST[minEdge.End] {
			// jeżeli wierzchołek końcowy nie jest w MST to dodaj krawędź do MST
			mst.AddEdge(minEdge.Start, minEdge.End, minEdge.Weight)
			inMST[minEdge.End] = true

			// dodaj wszystkie krawędzie wychodzące z wierzchołka końcowego do kolejki priorytetowej
			for _, edge := range inputGraph.GetAllEdgesFrom(minEdge.End) {
				if !inMST[edge.End] {
					heap.Push(pq, &edge)
				}
			}
		}
	}

	// sprawdzenie czy MST zawiera wszystkie wierzchołki
	for _, in := range inMST {
		if !in {
			return nil, 0
		}
	}

	return mst, 0
}
