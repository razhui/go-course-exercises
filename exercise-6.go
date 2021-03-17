//using iota, create 4 constants for the last 4 years. Print the constant values

package main

import (
	"fmt"
)

const (
	a = iota + 2018
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
