package main

import (
	"log"

	"tridentsk/streamcalc/utils"
)

func main() {

	pq := utils.NewPQueue(-1, "max")

	pq.Push(6)
	// log.Println(pq)
	pq.Push(5)
	// log.Println(pq)
	pq.Push(2)
	// log.Println(pq)
	pq.Push(1)
	// log.Println(pq)
	pq.Push(5)
	// log.Println(pq)
	pq.Push(4)
	// log.Println(pq)


	for i := 0; i < 6; i++ {
		log.Println(pq.Pop())
	}
}