package main

import "fmt"

func main() {
	maxSize := 5
	slice := make(map[string]int, maxSize)
	for i := range maxSize {
		slice[fmt.Sprintf("%d", i)] = i
	}
	chs := make(chan struct{}, 2)
	process2Ended := make(chan struct{})
	go func() {

		i := 0
		for {
			select {
			case <-process2Ended:
				chs <- struct{}{}
				return
			default:
				slice[fmt.Sprintf("%d", i%maxSize)] = i
				i++
			}
		}
	}()
	go func() {
		for i := 1000; i >= 0; i-- {
			res := slice[fmt.Sprintf("%d", i%maxSize)]
			fmt.Println(res)
		}
		process2Ended <- struct{}{}
		chs <- struct{}{}
	}()
	<-chs
	<-chs
}
