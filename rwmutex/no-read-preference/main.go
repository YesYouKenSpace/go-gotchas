package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	defer func() {
		//catch or finally
		if err := recover(); err != nil {
			//catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	mu := sync.RWMutex{}
	ch := make(chan struct{})
	go func() {
		mu.RLock()
		defer mu.RUnlock()
		ch <- struct{}{}
		fmt.Println("sleeping for 2ms")
		time.Sleep(2 * time.Millisecond)
		mu.RLock()
		defer mu.RUnlock()
		fmt.Printf("goroutine 1 ended\n")
		wg.Done()
	}()

	go func() {
		<-ch
		fmt.Println("go routine 2 started")
		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("goroutine 2 ended\n")
		wg.Done()
	}()
	wg.Wait()
}
