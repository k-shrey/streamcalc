package utils

import (
	"sync"
)

type Queue struct {
	sync.RWMutex
	*Deque
}

func newQueue(cap int) *Queue {
	return &Queue{
		Deque: NewDeque(cap),
	}
}

func (q *Queue) Enqueue(val float32) bool {
	q.Lock()
	defer q.Unlock()

	return q.PushBack(val)
}

func (q *Queue) Dequeue(val float32) interface{} {
	q.Lock()
	defer q.Unlock()

	return q.PopBack()
}
