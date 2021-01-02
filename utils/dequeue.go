package utils

import (
	"log"
	"os"
	"sync"
)

type Deque struct {
	sync.RWMutex
	container []float32
	capacity  int
}

func init() {
	log.SetOutput(os.Stdout)
}

func NewDeque(cap int) *Deque {
	return &Deque{
		capacity: cap,
	}
}

func (dq *Deque) Empty() bool {
	dq.RLock()
	defer dq.RUnlock()

	return len(dq.container) == 0
}

func (dq *Deque) Full() bool {
	dq.RLock()
	defer dq.RUnlock()

	return len(dq.container) == dq.capacity
}

func (dq *Deque) PushBack(val float32) bool {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		dq.container = append(dq.container, val)
		return true
	}

	log.Println("Failed to pushback ", val)
	return false
}

// helper function for efficient prepend
func prependValue(list []float32, val float32) []float32 {
	list = append(list, 0)
	copy(list[1:], list)
	list[0] = val

	return list
}

func (dq *Deque) PushFront(val float32) bool {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		dq.container = prependValue(dq.container, val)
		return true
	}

	log.Println("Failed to pushfront ", val)
	return false
}

// need to return nil on failure
func (dq *Deque) PopFront() interface{} {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		val := dq.container[0]
		dq.container = dq.container[1:]
		return val
	}

	log.Println("Failed to popfront ")
	return nil
}

// need to return nil on failure
func (dq *Deque) PopBack() interface{} {
	dq.Lock()
	defer dq.Unlock()

	if dq.capacity > 0 && dq.capacity > len(dq.container) {
		val := dq.container[0]
		dq.container = dq.container[:len(dq.container)-1]
		return val
	}

	log.Println("Failed to popback ")
	return nil
}

func (dq *Deque) Capacity() int {
	dq.RLock()
	defer dq.RUnlock()

	return dq.capacity
}
