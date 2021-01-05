package utils

import (
	"sync"
)

type Queue struct {
	sync.RWMutex
	*Deque
}

func NewQueue(cap int) Queue {
	return Queue{
		Deque: NewDeque(cap),
	}
}

func (q *Queue) Enqueue(val interface{}) bool {
	q.Lock()
	defer q.Unlock()

	return q.PushBack(val)
}

func (q *Queue) Dequeue(val interface{}) interface{} {
	q.Lock()
	defer q.Unlock()

	return q.PopFront()
}
