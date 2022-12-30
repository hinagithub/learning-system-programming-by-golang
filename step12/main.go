package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lestrrat/go-server-starter/listener"
)

func main() {

	//  Main11()
	// Main11_7_2()

	// signals := make(chan os.Signal, 1)
	// signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// s := <-signals
	// switch s {
	// case syscall.SIGINT:
	// 	fmt.Println("SITINT")
	// case syscall.SIGTERM:
	// 	fmt.Println("SIGTERM")
	// }

	// fmt.Println("Accept ctl + C for 10 second")
	// time.Sleep(time.Second * 10)
	// signal.Ignore(syscall.SIGINT, syscall.SIGHUP)
	// fmt.Println("Ignore ctl + C for 10 second")
	// time.Sleep(time.Second * 10)

	// signals := make(chan os.Signal, 1)
	// signal.Notify(signals, syscall.SIGINT,syscall.SIGTERM)
	// s:= <- signals

	// switch s {
	// case syscall.SIGINT:
	// 	fmt.Println(" ")
	// }

	// if len(os.Args) < 2 {
	// 	fmt.Printf("usage")
	// 	return
	// }
	// pid, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	// process, err := os.FindProcess(pid)
	// if err != nil {
	// 	panic(err)
	// }
	// process.Signal(os.Kill)
	// process.Kill()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid: %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])
	<-signals
	server.Shutdown(context.Background())

}
