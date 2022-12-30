package main
import (
	"fmt"
	 "time"
	 "math"
	 "context"
	 "os"
	 "os/signal"
	 "syscall"
)

func main () {

	Main4_3_1()
	
}

// goroutineサンプル
func Main4_1(){
	fmt.Println("start sub()")
	go sub()
	time.Sleep(2 * time.Second)
}
func sub(){
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}
func Main4_1_2(){
	fmt.Println("start sub()")
	// インラインで無名関数を作ってその場でgoroutineで実行　
	go func(){
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
func Main4_2(){
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func(){
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		done <- true
	}()
	<-done
	fmt.Println("all tasks are finished")
}


// for文とチャネル
func Main4_2_3(){
	pn :=primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}
func primeNumber()chan int{
	result := make(chan int) 
	go func() {
		result <- 2
		for i := 3; i < 10000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l + 1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

// Context
func Main4_2_5(){
	fmt.Println("start sub()")
	// 終了を受け取るための終了関数付きコンテキスト
	ctx, cancel := context.WithCancel(context.Background())
	// 終了時間を設定したり、タイムアウトの期限を設定するならcontext.WithDoneline()やcontext.WithTimeout()を使う
	go func (){
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		// 終了通知
		cancel()
	}()
	// 終了を待つ
	<- ctx.Done()
	fmt.Println("all tasks are finished")
}

// シグナル
func Main4_3_1(){
	// サイズが1より大きいチャネルを作成
	signals := make(chan os.Signal, 1)
	// SIGINT(CTRL+C)を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルが来るまで待つ
	fmt.Println("Waiting SIGINT (CTRL + C)")
	<- signals
	fmt.Println("SIGINT arrived")
}