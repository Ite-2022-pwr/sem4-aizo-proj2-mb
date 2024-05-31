package myGraph

import (
	"container/heap"
	"fmt"
	"projekt2/timeTrack"
	"time"
)

// Priority Queue for Edges
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

// Prim's algorithm to find the Minimum Spanning Tree (MST) of a graph
func Prim(inputGraph Graph, startVertex int, incidenceOrPredecessor bool) (mst Graph, elapsed int64) {
	startTime := time.Now()
	defer func() {
		elapsed = timeTrack.TimeTrack(startTime, "Prim")
	}()
	// Step 1: Initialize the MST based on the desired representation
	if incidenceOrPredecessor {
		mst = NewIncidenceMatrix()
	} else {
		mst = NewPredecessorList()
	}

	verticesCount := inputGraph.GetVertexCount()
	for i := 0; i < verticesCount; i++ {
		mst.AddVertex()
	}

	// Step 2: Create a priority queue to store the edges
	pq := &EdgePriorityQueue{}
	heap.Init(pq)

	// Step 3: Track the vertices that have been added to the MST
	inMST := make([]bool, verticesCount)
	inMST[startVertex] = true

	// Add all edges from the start vertex to the priority queue
	for _, edge := range inputGraph.GetAllEdgesFrom(startVertex) {
		heap.Push(pq, &edge)
	}

	// Step 4: Process the edges in the priority queue
	for pq.Len() > 0 {
		// Extract the edge with the minimum weight
		minEdge := heap.Pop(pq).(*Edge)
		fmt.Println(inMST)
		if !inMST[minEdge.End] {
			// If the end vertex is not in the MST, add the edge to the MST
			mst.AddEdge(minEdge.Start, minEdge.End, minEdge.Weight)
			inMST[minEdge.End] = true

			// Add all edges from the newly added vertex to the priority queue
			for _, edge := range inputGraph.GetAllEdgesFrom(minEdge.End) {
				if !inMST[edge.End] {
					heap.Push(pq, &edge)
				}
			}
		}
	}

	// Step 5: Ensure the resulting MST is valid
	for _, in := range inMST {
		if !in {
			return nil, 0
		}
	}

	return mst, 0
}
