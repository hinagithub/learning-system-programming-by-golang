package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {

	Main6_7_2()

}

// HTTPサーバ
func Main6_5_1() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
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
func Main6_5_2() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

}

// KeepAlice(サーバ側)

func Main6_6_1() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// Accept後のソケットで何度も応答を返すためにループ
			for {
				// タイムアウトを設定
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				//リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					//それ以外はエラーになる
					neterr, ok := err.(net.Error) // ダウンキャスト
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				//リクエストを表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello world\n"

				// レスポンスを書き込む
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          io.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}

// KeepAlive(クライアント側)
func Main6_6_2() {

	sendMessage := []string{
		"ASCII",
		"PROGRAMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	//　リトライ用にループで全体を囲う
	for {
		var err error
		//まだコネクションを張っていない　/ エラーでリトライ
		if conn == nil {
			//Dialから行ってconnを初期化
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessage[current]))
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		//サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// 結果を表示
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		current++
		if current == len(sendMessage) {
			break
		}
	}
	conn.Close()
}

// gzip
func Main6_7_1() {

	sendMessage := []string{
		"ASCII",
		"PROGRAMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	//　リトライ用にループで全体を囲う
	for {
		var err error
		//まだコネクションを張っていない　/ エラーでリトライ
		if conn == nil {
			//Dialから行ってconnを初期化
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessage[current]))
		request.Header.Set("Accept-Encoding", "gzip") // gzipを許可
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		//サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// 結果を表示
		dump, err := httputil.DumpResponse(response, false) // 2つ目の引数はbodyをダンプするかどうか。falseでbodyを無視する。
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		current++
		if current == len(sendMessage) {
			break
		}
	}
	conn.Close()
}

// クライアントはgzipを受け入れ可能か?
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func Main6_7_3() {

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}

}

// 1セッションの処理をする
func processSession(conn net.Conn) {
	// fmt.Printf("Accept: %d\n  conn.RemoteAddr()")
	// defer conn.Close()
	// for {
	// 	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // リクエストを読み込む
	// 	request, err := http.ReadRequest(bufio.NewReader(conn))
	// 	if err != nil {
	// 		neterr, ok := err.(net.Error)
	// 		if ok && neterr.Timeout() {
	// 			fmt.Println("Timeout")
	// 			break
	// 		} else if err == io.EOF {
	// 			break
	// 		}
	// 		panic(err)
	// 	}
	// 	dump, err := httputil.DumpRequest(request, true)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(dump))
	// 	// レスポンスを書き込む
	// 	response := http.Response{
	// 		StatusCode: 200,
	// 		ProtoMajor: 1,
	// 		ProtoMinor: 1,
	// 		Header:     make(http.Header),
	// 	}
	// 	if isGZipAcceptable(request) {
	// 		content := gzip.NewWriter(&buffer)
	// 		// コンテンツをgzipして転送
	// 		var buffer bytes.Buffer
	// 		writer := gzip, NewWriter(&buffer)
	// 		if isGZipAcceptable(request) {
	// 			content := "Hello World (gzipped)\n" // コンテンツをgzip化して転送
	// 			var buffer bytes.Buffer
	// 			writer := gzip.NewWriter(&buffer)
	// 			io.WriteString(writer, content)
	// 			writer.Close()
	// 			response.Body = io.NopCloser(&buffer)
	// 			response.ContentLength = int64(buffer.Len())
	// 			response.Header.Set("Content-Encoding", "gzip")
	// 		} else {
	// 			content := "Hello World\n"
	// 			response.Body = io.NopCloser(
	// 				strings.NewReader(content))
	// 			response.ContentLength = int64(len(content))
	// 		}
	// 		response.Write(conn)
	// 	}
	// }
}
