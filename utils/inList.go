package utils

func InListInt(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
