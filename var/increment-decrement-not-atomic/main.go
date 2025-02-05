package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	i := 0
	runs := 10000
	mu := sync.RWMutex{}
	mu.Lock()
	for range 100 {

		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				i++
			}
		}()
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				i--
			}
		}()
	}
	time.Sleep(1 * time.Second)
	mu.Unlock()
	time.Sleep(1 * time.Second)
	mu.Lock()
	fmt.Printf("i=%d\n", i)
}
