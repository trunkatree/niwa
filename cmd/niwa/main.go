package main

import (
	"fmt"
	"strconv"
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

	var i int = 123
	var i64 int64
	i64 = int64(i)      // 明示的な型変換が必要
	var b bool = i != 0 // boolへの変換には比較演算子を使う
	fmt.Println(i, i64, b)

	var istr string
	var iint int64
	istr = strconv.FormatInt(i64, 10)           // 文字列への変換はFormat系関数
	iint, err := strconv.ParseInt(istr, 10, 64) // 文字列からの変換はParse系関数
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(istr, iint)
}
