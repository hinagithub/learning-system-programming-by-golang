package main

import (
	"fmt"
)

// インターフェースを定義
type Talker interface {
	Talk()
}

// 構造体を宣言
type Greeter struct {
	name string
}

// 構造体はTalkerインターフェースで定義されているメソッドを持っている
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s \n", g.name)
}

func Main2() {
	// インターフェースの方を持つ変数を宣言
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
