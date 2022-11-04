package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// curl -G --data-urlencode "query=hello world" https://google.com
func main() {
	// クエリー文字列を作成
	values := url.Values{
		"query": {"hello world"},
	}
	// values.Encode() で文字列に変換
	resp, err := http.Get("https://google.com" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println(resp.Request.URL) // https://www.google.com/?query=hello+world
}
