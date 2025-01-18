package main

import "fmt"

func main() {
	var slice []int
	slice = nil
	// len(nil) is 0
	fmt.Printf("len(nil)=%d", len(slice)) // 0
}
