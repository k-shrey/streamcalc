package utils

import (
	"fmt"
	// "log"
	"sync"
)

type PQueue struct {
	sync.RWMutex
	container []float32
	compare   func(float32, float32) bool
	capacity  int
}

func (pq *PQueue) String() string{
	// `for idx, val := range pq.container {
	// 	log.Println(idx, ": ", val)
	// }`
	// fmt.Println("um????")
	// for idx, val := range pq.container {
	// 	log.Println(idx, ": ", val)
	// }
	return fmt.Sprintf("%#v", pq.container)
}
func NewPQueue(cap int, heaptype string) *PQueue {
	if heaptype == "max" {
		return &PQueue{
			compare:  maxComp,
			capacity: cap,
		}
	} else {
		return &PQueue{
			compare:  minComp,
			capacity: cap,
		}
	}

}

func maxComp(i, j float32) bool {
	return i > j
}

func minComp(i, j float32) bool {
	return i < j
}

func (pq *PQueue) Size() int {
	return len(pq.container)
}
func (pq *PQueue) Push(val float32) bool {
	pq.Lock()
	defer pq.Unlock()

	pq.container = append(pq.container, val)
	pq.heapifyUp(len(pq.container) - 1)

	return true
}

func (pq *PQueue) Pop() float32 {
	pq.Lock()
	defer pq.Unlock()

	top := pq.container[0]
	pq.container[0] = pq.container[len(pq.container)-1]
	pq.container = pq.container[:len(pq.container)-1]
	pq.heapifyDown(0)

	return top
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return 2*idx + 1
}

func right(idx int) int {
	return 2*idx + 2
}

func (pq *PQueue) isLeaf(idx int) bool {
	return idx >= len(pq.container)/2 && idx <= len(pq.container)
}
func (pq *PQueue) swap(i, j int) {
	pq.container[i], pq.container[j] = pq.container[j], pq.container[i]
}

func (pq *PQueue) heapifyUp(idx int) {
	for idx != 0 && pq.compare(pq.container[idx], pq.container[parent(idx)]) {
		// log.Println("swapping: ", pq.container[idx], pq.container[parent(idx)])
		pq.swap(idx, parent(idx))
		// log.Println("swapped: ", pq.container[idx], pq.container[parent(idx)])
		idx = parent(idx)
	}
}

func (pq *PQueue) heapifyDown(idx int) {

	// if pq.isLeaf(idx) {
	// 	return
	// }
	// log.Println("minhypg called with i = ", idx, " where val is: ", pq.container[idx])
	// log.Printf("%#v", pq.container)
	toSwap := idx
	left := left(idx)
	right := right(idx)
	size := len(pq.container)

	if left < size && pq.compare(pq.container[left], pq.container[idx]) {
		// log.Println(pq.container[left], " is less than ", pq.container[idx])
		toSwap = left
	}
	if right < size && pq.compare(pq.container[right], pq.container[toSwap]) {
		// log.Println(pq.container[right], " is less than ", pq.container[idx])
		toSwap = right
	}

	if toSwap != idx {
		pq.swap(toSwap, idx)
		pq.heapifyDown(toSwap)
	}

	return
}
