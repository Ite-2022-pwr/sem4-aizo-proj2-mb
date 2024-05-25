package utils

type Queue struct {
	Items []any
}

func NewQueue() *Queue {
	return &Queue{Items: make([]any, 0)}
}
func (q Queue) push(item any) {
	q.Items = append(q.Items, item)
}

func (q Queue) pop() any {
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}
