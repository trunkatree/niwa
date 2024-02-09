package main

import (
	"fmt"
)

var shadow bool

// 関数やメソッド、構造体などの型の定義の前に空行なしでコメントを書くと、ドキュメントとして利用される
func main() {
	n, _ := fmt.Println("Hello, world!", "...and", "Hello, world!")
	fmt.Println(n, "bytes written.")

	/*
	 * シャドーイング
	 * コードブロックの中で宣言された変数は、コードブロックの外の同名の変数に影響をおよぼさない
	 */
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
