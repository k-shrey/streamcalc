package utils

import (
	"sync"
)

// Expected data format
type Tick struct {
	Instrument string  `json:"instrument"`
	Price      float32 `json:"price"`
	Timestamp  int64   `json:"timestamp"`
}

// contains map[string]InstrumentData
type TickCache struct {
	sync.Map
}

// Holds data and statistics for an instrument
type InstrumentData struct {
	sync.RWMutex
	Data    Queue
	MinQ    Queue
	MaxQ    Queue
	Average float32
}

// expected json response for stats request
type Statistics struct {
	Average float32 `json:"avg"`
	Min     float32 `json:"min"`
	Max     float32 `json:"max"`
	Count   int     `json:"count"`
}

func ZeroIfEmpty(tick interface{}) float32 {
	if val, ok := tick.(*Tick); ok {
		return val.Price
	}
	return 0
}
