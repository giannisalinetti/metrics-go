package monitor

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// StatsManager manages a metrics interface and can be used with different
// kind of runtime data (memory, cpu, etc)
func StatsManager(mon Metrics, outFile string, printFreq int64, stop chan bool, wg *sync.WaitGroup) {
	// Set default collection time to 100 milisecongs
	var getFreq int64 = 100
	// But getFreq can't be higher than printfreq
	if getFreq > printFreq {
		getFreq = printFreq
	}

	// If WaitGroup is not nil add a goroutine instance to the delta
	if wg != nil {
		wg.Add(1)
	}

	// Configure logging output
	var fd *os.File
	var err error
	if outFile != "" {
		fd, err = os.Create(outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
	} else {
		fd = os.Stdout
	}
	logger := log.New(fd, "", log.LstdFlags)

	// Use a goroutine to update MemoryMonitor struct
	go mon.GetStats(stop, getFreq, wg)

	for {
		select {
		case <-stop:
			fmt.Println("Stopping stats printer.")
			//Since the data in the channel has been consumed we also send a new
			//signal to the GetStats goroutine
			stop <- true
			if wg != nil {
				wg.Done()
			}
			return
		default:
			time.Sleep(time.Duration(printFreq) * time.Millisecond)
			res, err := mon.PrintStats()
			if err != nil {
				log.Fatal(err)
			}
			logger.Println(res)
		}
	}
}
