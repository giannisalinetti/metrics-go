package monitor

import (
	"encoding/json"
	"fmt"
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
			if wg != nil {
				wg.Done()
			}
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
