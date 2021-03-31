package main //メインパッケージに関するプログラムと宣言, パッケージの宣言は１つのみ

import (
	"fmt" //標準パッケージのformatパッケージ
	"time"
)

/*
複数行のコメント
はこのようにつくる
*/

// 実際に実行されるのはmain関数の中身
func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())
}
