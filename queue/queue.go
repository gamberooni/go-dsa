package queue

import "errors"

const (
	ErrQueueIsFull  = QueueError("cannot enqueue a full queue")
	ErrQueueIsEmpty = QueueError("cannot dequeue an empty queue")
)

type QueueError string

func (q QueueError) Error() string {
	return string(q)
}

type Queue struct {
	Front    int
	Rear     int
	Elements []int
	Capacity int
}

func NewQueue(capacity int, elements []int) *Queue {
	q := Queue{}
	q.Front = -1
	q.Rear = -1
	q.Capacity = capacity
	if len(elements) > 0 {
		q.Elements = append(q.Elements, elements...)
		q.Front = 0
		q.Rear = len(elements) - 1
	}
	return &q
}

func (q Queue) Peek() int {
	if q.IsEmpty() {
		return 0
	}
	return q.Elements[q.Front]
}

func (q *Queue) Enqueue(element int) error {
	if q.IsFull() {
		return ErrQueueIsFull
	}
	q.Front = 0
	q.Rear += 1
	q.Elements = append(q.Elements, element)
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	switch {
	case q.IsEmpty():
		return 0, ErrQueueIsEmpty
	case len(q.Elements) > 1:
		dequeue := q.Elements[q.Front]
		q.Elements = q.Elements[1:]
		q.Rear -= 1
		return dequeue, nil
	case len(q.Elements) == 1:
		dequeue := q.Elements[q.Front]
		q.Elements = []int{}
		q.Front = -1
		q.Rear = -1
		return dequeue, nil
	}
	return 0, errors.New("unexpected error")
}

func (q Queue) IsFull() bool {
	return len(q.Elements) == q.Capacity
}

func (q Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}
