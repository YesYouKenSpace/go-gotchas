package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	atomicInt := atomic.Int64{}
	runs := 10000
	mu := sync.RWMutex{}
	mu.Lock()
	for range 100 {
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				atomicInt.Add(1)
			}
		}()
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				atomicInt.Add(-1)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	mu.Unlock()
	time.Sleep(1 * time.Second)
	mu.Lock()
	fmt.Printf("atomicInt=%d\n", atomicInt.Load())
}
