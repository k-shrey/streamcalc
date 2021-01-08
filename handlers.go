package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	. "tridentsk/streamcalc/utils"

	"github.com/gorilla/mux"
)

var tc TickCache

func TickHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var t Tick
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't parse tick", http.StatusBadRequest)
		return
	}

	if t.Timestamp < (time.Now().Unix() - 60) {
		http.Error(w, "Expired tick", http.StatusBadRequest)
		return
	}

	all, _ := tc.Load("ALL")
	ins, _ := tc.Load(t.Instrument)

	if ins == nil {
		ins = &InstrumentData{
			Data:    NewQueue(-1),
			MinQ:    NewQueue(-1),
			MaxQ:    NewQueue(-1),
			Average: 0,
		}
		tc.Store(t.Instrument, ins)
		go TidyWatcher(t.Instrument)
	}

	if all == nil {
		all = &InstrumentData{
			Data:    NewQueue(-1),
			MinQ:    NewQueue(-1),
			MaxQ:    NewQueue(-1),
			Average: 0,
		}
		tc.Store("ALL", all)
		go TidyWatcher("ALL")
	}

	var allData *InstrumentData = all.(*InstrumentData)
	var insData *InstrumentData = ins.(*InstrumentData)

	go calcStats(allData, &t)
	go calcStats(insData, &t)

	w.WriteHeader(http.StatusOK)
}

func calcStats(ins *InstrumentData, t *Tick) {
	ins.Lock()
	defer ins.Unlock()
	ins.Data.Enqueue(t)

	// update min
	for ins.MinQ.Size() > 0 && t.Price < ins.MinQ.Back().(*Tick).Price {
		ins.MinQ.PopBack()
	}
	ins.MinQ.Enqueue(t)

	// update max
	for ins.MaxQ.Size() > 0 && t.Price > ins.MaxQ.Back().(*Tick).Price {
		ins.MaxQ.PopBack()
	}
	ins.MaxQ.Enqueue(t)

	// update avg
	ins.Average = ins.Average + (t.Price-ins.Average)/float32(ins.Data.Size())

	// log.Println(ins.Average)
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ins, ok := vars["instrument"]
	if !ok {
		http.Error(w, "Missing instrument name", http.StatusBadRequest)
		return
	}

	insData, _ := tc.Load(ins)
	if insData == nil {
		http.Error(w, "No stats available for instrument", http.StatusBadRequest)
		return
	}

	var stats *InstrumentData = insData.(*InstrumentData)

	stats.RLock()
	payload := Statistics{
		Average: stats.Average,
		Min:     ZeroIfEmpty(stats.MinQ.Front()),
		Max:     ZeroIfEmpty(stats.MaxQ.Front()),
		Count:   stats.Data.Size(),
	}
	stats.RUnlock()

	jsonResp, _ := json.Marshal(payload)

	log.Printf("%#v\n", payload)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}

// monitor given instrument, remove timed out elements
func TidyWatcher(insName string) {
	insData, _ := tc.Load(insName)
	var ins *InstrumentData = insData.(*InstrumentData)

	for range time.Tick(time.Millisecond * 250) {
		ins.Lock()
		for ins.Data.Size() > 0 &&
			ins.Data.Front().(*Tick).Timestamp < (time.Now().Unix()-20) {

			value, ok := ins.Data.PopFront().(*Tick)

			// clean up MinQ
			for ok && ins.MinQ.Size() > 0 && value == ins.MinQ.Front().(*Tick) {
				ins.MinQ.PopFront()
			}

			// clean up MaxQ
			for ok && ins.MaxQ.Size() > 0 && value == ins.MaxQ.Front().(*Tick) {
				ins.MaxQ.PopFront()
			}

			// update average
			if ins.Data.Size() > 0 {
				ins.Average = ins.Average -
					(value.Price-ins.Average)/float32(ins.Data.Size())
			} else {
				ins.Average = 0
			}

		}
		ins.Unlock()
	}
}
