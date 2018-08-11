package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/giannisalinetti/metrics-go/pkg/monitor"
)

// businessLogic is just a dummy function whose purpose is to print on screen
// something and return.
func businessLogic() error {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Doing some crazy stuff to entertain you...")
	}
	return nil
}

func main() {
	freqFlag := flag.Int64("f", 1000, "Print frequency")
	flag.Parse()

	//Define a stop channel to handle the StatsManager
	stopCh := make(chan bool)
	defer close(stopCh)

	//Using sync.WaitGroup to handle goroutines closing
	wg := &sync.WaitGroup{}

	//Allocate new MemoryMonitor
	mmon := monitor.NewMemoryMonitor()

	//StatsManager should run in the background, letting the main program
	//logic do something else.
	wg.Add(1)
	go monitor.StatsManager(mmon, *freqFlag, stopCh, wg)

	//Businesslogic goes here, running in foreground
	err := businessLogic()
	if err != nil {
		log.Fatal(err)
	}

	//Send a signal to cascading stop the goroutines
	stopCh <- true
	wg.Wait()
}
