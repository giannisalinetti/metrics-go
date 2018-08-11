package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

// MemoryMonitor is a struct for memory metrics collection
type MemoryMonitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64
	NumGC        uint32
	NumGoroutine int
	mutex        sync.RWMutex
}

// Metrics define a common interface for metrics
type Metrics interface {
	GetStats(ch chan bool, freq int64, wg *sync.WaitGroup)
	PrintStats() (string, error)
}

// NewMemoryMonitor return a *MemoryMonitor pointer
func NewMemoryMonitor() *MemoryMonitor {
	//Return an empty Monitor struct
	return &MemoryMonitor{}
}

// GetStats collects stats using the runtime package. It is designed to be
// executed as a goroutine.
func (m *MemoryMonitor) GetStats(ch chan bool, freq int64, wg *sync.WaitGroup) {

	//Load data into runtime.MemStats struct
	rtm := runtime.MemStats{}

	fmt.Println("Starting memory stats collector.")

	for {
		select {
		case <-ch:
			fmt.Println("Stopping memory stats collector.")
			wg.Done()
			return
		default:
			time.Sleep(time.Duration(freq) * time.Millisecond)
			m.mutex.Lock()
			runtime.ReadMemStats(&rtm)

			//Put data into the Monitor struct
			m.NumGoroutine = runtime.NumGoroutine()
			m.Alloc = rtm.Alloc
			m.TotalAlloc = rtm.TotalAlloc
			m.Sys = rtm.Sys
			m.Mallocs = rtm.Mallocs
			m.Frees = rtm.Frees
			m.LiveObjects = m.Mallocs - m.Frees
			m.PauseTotalNs = rtm.PauseTotalNs
			m.NumGC = rtm.NumGC
			m.mutex.Unlock()
		}
	}
}

// PrintStats returns a json string with the data loaded in MemoryMonitor
func (m *MemoryMonitor) PrintStats() (string, error) {
	output, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// StatsManager manages a metrics interface and can be used with different
// kind of runtime data (memory, cpu, etc)
func StatsManager(mon Metrics, printFreq int64, stop chan bool, wg *sync.WaitGroup) {
	//Set default collection time to 100 milisecongs
	var getFreq int64 = 100

	//getFreq can't be higher than printfreq
	if getFreq > printFreq {
		getFreq = printFreq
	}

	//Use a goroutine to update Monitor every 100 milliseconds
	wg.Add(1)
	go mon.GetStats(stop, getFreq, wg)

	for {
		select {
		case <-stop:
			fmt.Println("Stopping stats printer.")
			//Since the data in the channel has been consumed we also send a new
			//signal to the GetStats goroutine
			stop <- true
			wg.Done()
			return
		default:
			time.Sleep(time.Duration(printFreq) * time.Millisecond)
			res, err := mon.PrintStats()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
	}
}
