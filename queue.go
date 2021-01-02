package dequeue

import (
	// "fmt"
	"sync"
)

type Dequeue struct {
	sync.RWMutex
	container []float32
	capacity int
}

func newDequeue(cap int) *Dequeue {
	return &Dequeue{
		capacity: cap,
	}
}
func (dq *Dequeue) Empty() bool {
	dq.RLock()
	defer dq.RUnlock()

	if len(dq.container) == 0 {
		return true
	}
	return false
}

func (dq *Dequeue) Full() bool {
	dq.RLock()
	defer dq.RUnlock()

	if len(dq.container) == dq.capacity {
		return true
	}
	return false
}

func (dq *Dequeue) PushBack(val float32) bool {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		dq.container = append(dq.container, val)
		return true
	}

	return false
}

// helper function for efficient prepend
func prependValue(list []float32, val float32) []float32 {
	list = append(list, 0)
	copy(list[1:], list)
	list[0] = val

	return list
}

func (dq *Dequeue) PushFront(val float32) bool {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		dq.container = prependValue(dq.container, val)
		return true
	}
	return false
}

// using interface{}, need to be able to return nil on failure
func (dq *Dequeue) PopFront() interface{} {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		val := dq.container[0]
		dq.container = dq.container[1:]
		return val
	}

	return nil
}

func (dq *Dequeue) PopBack() interface{} {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		val := dq.container[0]
		dq.container = dq.container[:len(dq.container) - 1]
		return val
	}

	return nil
}

func (dq *Dequeue) Capacity() int {
	dq.RLock()
	defer dq.RUnlock()

	return dq.capacity
}




