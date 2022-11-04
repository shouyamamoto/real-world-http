package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// curl -T main.go "Content-Type: text/plain" http://localhost:18888
func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	// Postには、送信する内容をテキスト化するのではなく、io.Reader型で渡す
	// file は os.File型であるが、io.Readerインターフェースを満たしているため引数に渡すことができる
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)

	// テキストをhttp.Postに渡す場合は以下のように文字列をio.Readerインターフェース化する
	// io.Readerインターフェース化させるには、 bytes.Bufferやstrings.NewReaderを使う
	reader := strings.NewReader("テキスト")
	resp, err = http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
