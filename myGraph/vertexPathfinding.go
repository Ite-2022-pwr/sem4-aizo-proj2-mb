package myGraph

type VertexPathfinding struct {
	Index         int
	Predecessor   int
	WeightToStart int
	Visited       bool
}

func SortByWeightToStartQS(list []VertexPathfinding) []VertexPathfinding {
	if len(list) < 2 {
		return list
	}
	left, right := 0, len(list)-1
	pivot := len(list) / 2
	list[pivot], list[right] = list[right], list[pivot]
	for i := range list {
		if list[i].WeightToStart < list[right].WeightToStart {
			list[i], list[left] = list[left], list[i]
			left++
		}
	}
	list[left], list[right] = list[right], list[left]
	SortByWeightToStartQS(list[:left])
	SortByWeightToStartQS(list[left+1:])
	return list
}

func FindByIndex(list []VertexPathfinding, index int) (out *VertexPathfinding) {
	for i := 0; i < len(list); i++ {
		if list[i].Index == index {
			return &list[i]
		}
	}
	return &VertexPathfinding{}
}
