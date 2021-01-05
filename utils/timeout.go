package utils

import "time"

func TidyWatcher(ins string) {
	// monitor each list in TickCache, remove timed out elements
	for range time.Tick(time.Millisecond * 250) {

	}
}
