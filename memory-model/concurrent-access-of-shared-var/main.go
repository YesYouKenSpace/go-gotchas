package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	num := 1.0
	runs := 10000
	mu := sync.RWMutex{}
	mu.Lock()
	for range 100 {
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				if float64(int(num/2)) == num/2 {
					num /= 2
				}
			}
		}()
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			for range runs {
				num++
			}
		}()
	}
	time.Sleep(1 * time.Second)
	mu.Unlock()
	time.Sleep(1 * time.Second)
	mu.Lock()
	// Is the number always a whole number?
	fmt.Printf("num=%f\n", num)
}
