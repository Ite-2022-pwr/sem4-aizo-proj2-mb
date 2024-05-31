package myGraph

import (
	"errors"
	"projekt2/utils"
)

// Kruskal algorithm to find the Minimum Spanning Tree (MST) of a graph
func Kruskal(inputGraph Graph, outputIncidenceOrPredecessor bool) (mst Graph, err error) {
	// Step 1: Sort all edges in the input graph by weight in non-decreasing order
	sortedEdges := SortEdgesListQS(inputGraph.GetAllEdges())

	// Step 2: Initialize the MST based on the desired representation
	if outputIncidenceOrPredecessor {
		mst = NewIncidenceMatrix()
	} else {
		mst = NewPredecessorList()
	}

	// Add all vertices to the MST graph
	for i := 0; i < inputGraph.GetVertexCount(); i++ {
		mst.AddVertex()
	}

	// Step 3: Create a map to track the connected components of each vertex
	verticesMap := make(map[int]int)
	for i := 0; i < mst.GetVertexCount(); i++ {
		verticesMap[i] = i
	}

	// List to keep track of connected vertices
	connected := make([]int, 0)

	// Step 4: Iterate over sorted edges and add them to the MST if they don't form a cycle
	for _, edge := range sortedEdges {
		// Stop if we have enough edges to form a spanning tree
		if len(connected) == mst.GetVertexCount() {
			break
		}

		start := edge.Start
		end := edge.End

		// Check if the current edge forms a cycle
		if start != end && !mst.IsAdjacent(start, end) && verticesMap[start] != verticesMap[end] {
			// Add edge to the MST
			mst.AddEdge(start, end, edge.Weight)

			// Update the connected components map
			oldComponent := verticesMap[start]
			newComponent := verticesMap[end]
			for vertex, component := range verticesMap {
				if component == oldComponent {
					verticesMap[vertex] = newComponent
				}
			}

			// Add vertices to the connected list
			if !utils.InListInt(connected, end) {
				connected = append(connected, end)
			}
			if !utils.InListInt(connected, start) {
				connected = append(connected, start)
			}
		}
	}

	// Step 5: Ensure the resulting MST is valid
	if len(connected) != mst.GetVertexCount() && mst.GetEdgeCount() != mst.GetVertexCount()-1 {
		return nil, errors.New("unable to form a valid MST")
	}

	return mst, nil
}
