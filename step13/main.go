package main

import (
	"fmt"
	"time"
)

func main() {

	Main13_1()

}

func sub1(c int) {
	fmt.Println("Share by arguments:", c*c)
}

func Main13_1() {
	go sub1(10)
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	time.Sleep(time.Second)
}
