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
		// 1. acquire the read lock
		fmt.Println("goroutine A started")
		mu.RLock()
		fmt.Println("goroutine A acquired the read lock")
		defer mu.RUnlock()
		// 3. enable the other goroutine to attempt to acquire the write lock
		fmt.Println("goroutine A sending signal to goroutine B")
		ch <- struct{}{}
		fmt.Println("goroutine A sleeping for 2ms")
		time.Sleep(2 * time.Millisecond)
		// 5. attempt to acquire another read lock -- this would succeed if there was read preference and eventually unblock goroutine 2
		fmt.Println("goroutine A attempting to acquire another read lock -- if there was read preference, this would succeed")
		mu.RLock()
		defer mu.RUnlock()
		fmt.Printf("goroutine A ended\n")
		wg.Done()
	}()

	go func() {
		fmt.Println("goroutine B started")
		// 2. wait for the read lock to be acquired
		fmt.Println("goroutine B waiting for goroutine A to acquire the read lock")
		<-ch
		// 4. attempt to acquire the write lock -- we will be stuck here forever
		fmt.Println("goroutine B attempting to acquire the write lock")
		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("goroutine B ended\n")
		wg.Done()
	}()
	wg.Wait()
}
