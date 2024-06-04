package myGraph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// odczytywanie z pliku
func ReadFromFile(filename string, directed bool, useIncidenceMatrix bool) (Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var graph Graph
	if useIncidenceMatrix {
		graph = NewIncidenceMatrix()
	} else {
		graph = NewPredecessorList()
	}
	graph.SetDirected(directed)

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, fmt.Errorf("file is empty or invalid format")
	}
	firstLine := strings.Fields(scanner.Text())
	if len(firstLine) != 2 {
		return nil, fmt.Errorf("invalid first line format")
	}
	vertices, err := strconv.Atoi(firstLine[0])
	if err != nil {
		return nil, fmt.Errorf("invalid vertex count")
	}
	edges, err := strconv.Atoi(firstLine[1])
	if err != nil {
		return nil, fmt.Errorf("invalid edge count")
	}

	for i := 0; i < vertices; i++ {
		graph.AddVertex()
	}

	for i := 0; i < edges; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected end of file")
		}
		line := strings.Fields(scanner.Text())
		if len(line) != 3 {
			return nil, fmt.Errorf("invalid edge format on line %d", i+2)
		}
		start, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start vertex on line %d", i+2)
		}
		end, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, fmt.Errorf("invalid end vertex on line %d", i+2)
		}
		weight, err := strconv.Atoi(line[2])
		if err != nil {
			return nil, fmt.Errorf("invalid weight on line %d", i+2)
		}
		graph.AddEdge(start, end, weight)
	}

	return graph, nil
}
