//Print a number in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%d\n%b\n%#x", x, x, x)
}
