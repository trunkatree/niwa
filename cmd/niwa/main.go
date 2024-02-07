package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello, world!", "...and", "Hello, world!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, "bytes written.")
}
