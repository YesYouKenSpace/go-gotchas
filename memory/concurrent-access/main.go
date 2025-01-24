package main

import (
	"fmt"
	"time"
)

func main() {
	num := 1.0
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				if num == float64(int(num)) {
					dividedBy2 := num / 2
					rounded := float64(int(num / 2))
					if dividedBy2 == rounded {
						num = num / 2
					}
				}
			}
		}
	}()

	go func() {

		for {
			select {
			case <-ch:
				return
			default:
				num++
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				fmt.Printf("num=%f\n", num)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)
	// Is the number always a whole number?
	fmt.Printf("num=%f\n", num)
}
