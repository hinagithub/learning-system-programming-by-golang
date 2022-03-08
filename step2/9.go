package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
)

// フォーマット
func Main9_1() {
	fmt.Fprintf(os.Stdout, "d =「%d」= 1, s = [%s] = あああ, f = [%f] = 3.14", 1, "あああ", 3.14)
}

// CSV
func Main9_2() {
	file, err := os.Create("tmp/text.csv")
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(file)
	w.Comma = ','
	w.Write([]string{"first_name", "last_name", "username"})
	w.Write([]string{"mail", "female", "others"})
	w.Write([]string{"post_code", "prefecture", "city"})
	w.Flush()
}

func handler9(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// json元データ
	source := map[string]string{
		"Hello": "World",
	}

	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString(source["Hello"])
	buffer.Flush()

	io.WriteString(w, source["Hello"])
}

// gzip -> JSON -> Stdout
func Main9_3() {
	http.HandleFunc("/", handler9)
	http.ListenAndServe(":8081", nil)
}
