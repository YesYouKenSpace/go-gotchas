package main

import "fmt"

func print(i int) {
	fmt.Printf("print=%d\n", i)
}

func printByAddress(j *int) {
	fmt.Printf("printByAddress=%d\n", *j)
}

func main() {
	i := 0
	defer print(i)
	defer printByAddress(&i)
	i++
	print(i)
	printByAddress(&i)
}

// Here is the output of the program:
// print=2
// printByAddress=2
// printByAddress=2
// print=1

// This demonstrates that deferred functions are stacked, meaning they are executed LIFO (last in first out).
// The deferred function printByAddress is executed first, then print.

// Another behaviour demonstrated here is that the value of i is captured by value when the defer statement is executed.
// Think of it as the function and its arguments are pushed into the stack immediately when the defer statement is executed.
