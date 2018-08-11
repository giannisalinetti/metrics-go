package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64
	NumGC        uint32
	NumGoroutine int
}

func NewMonitor(duration int) {
	var m Monitor
	var rtm runtime.MemStats
	var interval = time.Duration(duration) * time.Second
	for {
		<-time.After(interval)

		//Read full mem stats
		runtime.ReadMemStats(&rtm)

		//Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		//Misc memory stats
		m.Alloc = rtm.Alloc
		m.TotalAlloc = rtm.TotalAlloc
		m.Sys = rtm.Sys
		m.Mallocs = rtm.Mallocs
		m.Frees = rtm.Frees

		//Live objects = Malloc - Frees
		m.LiveObjects = m.Mallocs - m.Frees

		//GC Stats
		m.PauseTotalNs = rtm.PauseTotalNs
		m.NumGC = rtm.NumGC

		//Encode to json and print
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
	}
}

func main() {
	NewMonitor(1)
}