package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unsafeInt := 0
	runs := 10000
	mu := sync.RWMutex{}
	mu.Lock()
	for range 100 {
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				unsafeInt++
			}
		}()
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				unsafeInt--
			}
		}()
	}
	time.Sleep(1 * time.Second)
	mu.Unlock()
	time.Sleep(1 * time.Second)
	mu.Lock()
	fmt.Printf("unsafeInt=%d\n", unsafeInt)
}
