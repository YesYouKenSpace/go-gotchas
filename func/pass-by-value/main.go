package main

import "fmt"

func pointerFn(e *int) {
	*e = 10
	fmt.Printf("at the end of pointerFn: %v\n", *e)
}

func valueFn(e int) {
	e = 10
	fmt.Printf("at the end of valueFn: %v\n", e)
}

// slices are essentially pointers to the start of the array
func sliceFn(e []int) {
	e[0] = 10
	fmt.Printf("at the end of sliceFn: %v\n", e)
}

// structs are copied in entirety
func structFn(e struct {
	Name         string
	WillIPersist *string
}) {
	e.Name = "John"
	*e.WillIPersist = "yes"
	fmt.Printf("at the end of structFn: %#v\n", e)
	fmt.Printf("WillIPersist: %v\n", *e.WillIPersist)
}

func main() {
	fmt.Println("# pointer function")

	a := 5
	fmt.Printf("before pointerFn: %v\n", a)
	pointerFn(&a)
	fmt.Printf("after pointerFn: %v\n", a)

	fmt.Println("")
	fmt.Println("# value function")

	b := 5
	fmt.Printf("before valueFn: %v\n", b)
	valueFn(b)
	fmt.Printf("after valueFn: %v\n", b)

	fmt.Println("")
	fmt.Println("# slice function")

	c := []int{5, 6, 7}
	fmt.Printf("before sliceFn: %v\n", c)
	sliceFn(c)
	fmt.Printf("after sliceFn: %v\n", c)

	fmt.Println("")
	fmt.Println("# struct function")
	d := struct {
		Name         string
		WillIPersist *string
	}{
		Name:         "Jane",
		WillIPersist: new(string),
	}
	*d.WillIPersist = "let's see"
	fmt.Printf("before structFn: %+v\n", d)
	fmt.Printf("WillIPersist: %v\n", *d.WillIPersist)
	structFn(d)
	fmt.Printf("after structFn: %+v\n", d)
	fmt.Printf("WillIPersist: %v\n", *d.WillIPersist)
}
