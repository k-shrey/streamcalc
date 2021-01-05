package utils

import (
	"sync"
	// . "tridentsk/streamcalc/utils"
)

type Tick struct {
	Instrument string  `json:"instrument"`
	Price      float32 `json:"price"`
	Timestamp  int64   `json:"timestamp"`
}

type TickCache struct {
	// one queue for each instrument, map[string][]Tick
	sync.Map
}

type InstrumentData struct {
	sync.RWMutex
	Data    Queue
	MinQ    Queue
	MaxQ    Queue
	Average float32
}

type Statistics struct {
	Average float32 `json:"avg"`
	Min     float32 `json:"min"`
	Max     float32 `json:"max"`
	Count   int     `json:"count"`
}

// type Average struct {
// 	sync.Map
// }

// type Min struct {
// 	sync.Map
// }

// type Max struct {
// 	sync.Map
// }
