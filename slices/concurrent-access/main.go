package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
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
				slice[i%len(slice)] = i
				i++
			}
		}
	}()
	go func() {
		for i := 1000; i >= 0; i-- {
			res := slice[i%len(slice)]
			fmt.Println(res)
		}
		process2Ended <- struct{}{}
		chs <- struct{}{}
	}()
	<-chs
	<-chs
}
