package main

import (
	"fmt"
	"sync"
)

func main() {
	// pool作成
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("manually added: 1")
	pool.Put("manually added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	// manually added: 1
	// manually added: 2
	// created: 1
}
