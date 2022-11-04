package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// curl -X DELETE http://localhost:18888
// httpモジュールまたはhttp.Client構造体のメソッドは
// GET, HEAD, POST のみメソッドをサポートしている。
// それ以外のメソッドを使用したい場合には、http.Request構造体のオブジェクトを使う必要がある。
func main() {
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
