package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// ファイル作成
func Main4() {
	file, err := os.Create("tmp/test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}

// 標準出力
func Main4_2() {
	os.Stdout.Write([]byte("os.Stdout example\nコンソールに出力\n"))
}

// 書かれた内容を記憶しておくバッファ(1)
func Main4_3() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// 書かれた内容を記憶しておくバッファ(2)
func Main4_4() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}

// インターネットアクセス
func Main4_5() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

// レスポンス
func Main4_5_2() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
}

// デコレータ
func Main4_6() {
	file, err := os.Create("tmp/multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}

// zip
func Main4_6_2() {
	file, err := os.Create("tmp/text.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "text.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}

func Main4_6_3() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
	buffer.WriteString("1\n")
	buffer.WriteString("2\n")
	buffer.Flush()
	buffer.WriteString("3\n")
}

// Fprintfフォーマット
func Main4_7() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}

// JSONフォーマット
func Main4_7_2() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
		"1":       "サンプル",
		"日本語":     "japanese",
	})
}
