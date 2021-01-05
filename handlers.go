package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "sync"

	// "time"
	// utils "tridentsk/streamcalc/utils"
	// . "tridentsk/streamcalc/types"
	. "tridentsk/streamcalc/utils"
)

var tc TickCache

func StoreValueHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var t Tick
	err = json.Unmarshal(body, &t)
	// if t.Timestamp < (time.Now().Unix() - 60) {
	// 	// expired val
	// 	log.Println("got an expired value")
	// 	w.WriteHeader(201)
	// }

	all, _ := tc.Load("All")
	arr, _ := tc.Load(t.Instrument)

	if all == nil {
		all = &InstrumentData{
			Data:    NewQueue(-1),
			MinQ:    NewQueue(-1),
			MaxQ:    NewQueue(-1),
			Average: 0,
		}
		tc.Store("All", all)
	}
	if arr == nil {
		arr = &InstrumentData{
			Data:    NewQueue(-1),
			MinQ:    NewQueue(-1),
			MaxQ:    NewQueue(-1),
			Average: 0,
		}
		tc.Store(t.Instrument, arr)
	}

	var allData *InstrumentData = all.(*InstrumentData)
	var arrData *InstrumentData = arr.(*InstrumentData)

	// var wg sync.WaitGroup
	// wg.Add(2)
	// go calcStats(allData, &t, &wg)
	// go calcStats(arrData, &t, &wg)
	// wg.Wait()
	go calcStats(allData, t)
	go calcStats(arrData, t)

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

// func calcStats(ins *InstrumentData, t *Tick, wg *sync.WaitGroup) {
func calcStats(ins *InstrumentData, t Tick) {

	ins.Lock()
	defer ins.Unlock()
	// defer wg.Done()
	ins.Data.Enqueue(t)

	// fix MinQ
	for ins.MinQ.Size() > 0 && t.Price < ins.MinQ.Back().(Tick).Price {
		ins.MinQ.PopBack()
	}
	ins.MinQ.Enqueue(t)

	// fix maxQ
	for ins.MaxQ.Size() > 0 && t.Price > ins.MaxQ.Back().(Tick).Price {
		ins.MaxQ.PopBack()
	}
	ins.MaxQ.Enqueue(t)

	ins.Average = ins.Average + (t.Price-ins.Average)/float32(ins.Data.Size())
	// log.Println(ins.Average)
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ins, ok := vars["instrument"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	insData, _ := tc.Load(ins)
	if insData == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var stats *InstrumentData = insData.(*InstrumentData)

	stats.RLock()
	payload := Statistics{
		Average: stats.Average,
		Min:     stats.MinQ.Front().(Tick).Price,
		Max:     stats.MaxQ.Front().(Tick).Price,
		Count:   stats.Data.Size(),
	}
	stats.RUnlock()

	jsonResp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}
