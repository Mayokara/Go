package main

import "fmt"

// i5 := 500
var i5 int = 500

/*基本的には型指定を行った変数定義を行う。
Goは静的片付け言語
*/

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100 //整数型
	fmt.Println(i)

	var s string = "Hello Go" //文字列型
	fmt.Println(s)

	var t, f bool = true, false //２つ同時に宣言も可能
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数の型だけ指定して値を格納しないでみる
	var i3 int    // 0が格納されてる
	var s3 string //空文字が格納されてる
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	//i4 := 500 //再定義はできない
	//fmt.Println(i4)

	// i4 = "Hello" //異なる型では更新できない
	// fmt.Println(i4)
	fmt.Println(i5)

	outer()
	//fmt.Println(s4) //s4はouter関数内のみで有効

	var s5 string = "Not Use" //Goは定義された変数はプログラム上で必ず使う必要がある
	fmt.Println(s5)
}
