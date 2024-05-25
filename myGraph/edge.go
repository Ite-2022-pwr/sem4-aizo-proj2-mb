package myGraph

type Edge struct {
	Start  int
	End    int
	Weight int
}

func (e *Edge) isInList(list []Edge, directed bool) bool {
	for i := 0; i < len(list); i++ {
		if directed {
			if list[i].Start == e.Start && list[i].End == e.End {
				return true
			}
		} else {
			if (list[i].Start == e.Start && list[i].End == e.End) || (list[i].Start == e.End && list[i].End == e.Start) {
				return true
			}
		}
	}
	return false
}

func SortEdgesListQS(edges []Edge) []Edge {
	if len(edges) < 2 {
		return edges
	}
	left, right := 0, len(edges)-1
	pivot := len(edges) / 2
	edges[pivot], edges[right] = edges[right], edges[pivot]
	for i := range edges {
		if edges[i].Weight < edges[right].Weight {
			edges[i], edges[left] = edges[left], edges[i]
			left++
		}
	}
	edges[left], edges[right] = edges[right], edges[left]
	SortEdgesListQS(edges[:left])
	SortEdgesListQS(edges[left+1:])
	return edges
}
