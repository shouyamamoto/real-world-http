package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// curl https://google.com
func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body)) // レスポンスのBodyをUTF-8文字列に変換して出力
	fmt.Println(resp.Status)  // レスポンスのステータスを出力
}
