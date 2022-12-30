package main
import (
	"fmt"
	"io"
	 "io/ioutil"
	 "strings"
	 "bytes"
	 "os"
	 "net"
)

func main () {
	Main2_3_2()
	Main2_4_3()
	Main2_4_4()
	Main2_4_5()
}

// 一般的なReader 
func Main2_3_2(){
	var reader io.Reader =strings.NewReader("テストデータ")
	var readCloser io.ReadCloser = ioutil.NopCloser(reader)
	fmt.Println(readCloser)
}

// bytes.Buffer
func Main2_4_3(){
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// strings.Buffer
func Main2_4_4(){
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}

// インターネットアクセス
func Main2_4_5(){
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}