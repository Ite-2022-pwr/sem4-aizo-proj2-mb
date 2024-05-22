package myGraph

type Edge struct {
	start  int
	end    int
	weight int
}

func (e *Edge) isInList(list []Edge, directed bool) bool {
	for i := 0; i < len(list); i++ {
		if directed {
			if list[i].start == e.start && list[i].end == e.end {
				return true
			}
		} else {
			if (list[i].start == e.start && list[i].end == e.end) || (list[i].start == e.end && list[i].end == e.start) {
				return true
			}
		}
	}
	return false
}
