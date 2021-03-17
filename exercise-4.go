//assign an int to a variable
//prints that it into decimal, binary, and hex
//shifts the bits of that int over 1 position to the left, and assigns that to a variable
//prints that variable into decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	a := 143
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\n%b\n%#x", b, b, b)
}
