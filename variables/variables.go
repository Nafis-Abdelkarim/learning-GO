package main

import "fmt"

// Declare a single variable
var a int

// Declare a group of vars
var (
	b bool
	c float32
	d string
)

func main() {
	a = 42
	b, c = true, 32.0
	d = "string"
	fmt.Println(a, b, c, d)
}
