package monitor

import (
	"sync"
)

// Metrics define a common interface for metrics
type Metrics interface {
	GetStats(ch chan bool, freq int64, wg *sync.WaitGroup)
	PrintStats() (string, error)
}
