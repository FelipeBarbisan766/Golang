package domains

import "fmt"

// Declarar Struct
type QueueArray struct {
	Items []int
}

func (q *QueueArray) Enqueue(data int) {
	q.Items = append(q.Items, data)
}

func (q *QueueArray) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *QueueArray) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	removed := q.Items[0]
	q.Items = q.Items[1:]
	return removed
}

func (q *QueueArray) Show() {
	for _, item := range q.Items {
		fmt.Println(item)
	}
	fmt.Println()
}

func (q *QueueArray) Size() int {
	return len(q.Items)
}
