package main

import (
	"fmt"
)

var shadow bool

func main() {
	n, _ := fmt.Println("Hello, world!", "...and", "Hello, world!")
	fmt.Println(n, "bytes written.")

	shadow = true
	// shadow = false
	x := 0
	if shadow {
		x := 1
		fmt.Println("inside of if: x =", x)
	} else {
		x = 1
		fmt.Println("inside of if: x =", x)
	}
	fmt.Println("outside of if: x =", x)
}
