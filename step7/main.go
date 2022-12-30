package main
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	// "io"
	// "time"
)

func main () {

	Main6_5_2()
	
}

// HTTPサーバ
func Main6_5_1(){
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err) 
	}
	fmt.Println("Server is running at localhost:8888")
	for{

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func (){
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil{
				panic(err) 
			}
			dump, err := httputil.DumpRequest(request, true)
			if(err != nil){
				panic(err)
			}
			fmt.Println(string(dump))
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello World  \n")),
			}
			response.Write(conn)
			conn.Close()
	
		}()
	}
}

// HTTPクライアント
func Main6_5_2(){
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil{
		panic(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err !=nil{
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err !=  nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(dump))

}

// func Main6_6_1

func Main6_6_1(){
	// defer conn.Close()
	// fmt.Printf("Accept %v\n")
	// // Accept後のソケットで何度も応答を返すためにループ　
	// for {
	// 	// タイムアウトを設定
	// 	conn.setReadDeadline(time, Now().Add(S * time.Second)) 
	// 	// リクエストを読み込む
	// 	request, err := http.ReadRequest(bufio.NewReader(conn))
	// 	if err != nil{

	// 		// タイムアウトもしくはソケットクローズ時は終了
	// 		neterr, ok := err.(net.Error)// ダウンキャスト
	// 		if ok {}
	// 			}

	// }
}


