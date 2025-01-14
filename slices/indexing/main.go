package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	// This will print [1 2 3 4 5]
	fmt.Printf("s=%+v\n", s)

	// This will print []
	fmt.Printf("s[0:0]=%+v\n", s[0:0])

	// This will print [1 2 3 4 5]
	fmt.Printf("s[0:5]=%+v\n", s[0:5])

	// This will print [1]
	fmt.Printf("s[0:1]=%+v\n", s[0:1])

	// Negative indexing is not allowed
	// fmt.Printf("s[0:1]=%+v\n", s[:-1])

	// Printing in reverse order is not allowed
	// fmt.Printf("s[len(s)-2:len(s)-3]=%+v\n", s[len(s)-2:len(s)-3])
}
