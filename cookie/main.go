package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	// Jarはcookiejarを指定する。
	// Jarは関連するcookieを各リクエストに挿入するために使用され、送信するリクエストにcookieを挿入し
	// 受信するレスポンスのcookieの値で更新される。
	// Jarはクライアントが従うリダイレクトごとにJarが参照される。
	// Jarがnilの場合、cookieはリクエストで明示的に設定された場合のみ送信される。
	client := http.Client{
		Jar: jar,
	}

	// 初回アクセスでcookieを受信し、2回目以降のアクセスでcookieをサーバーに対して送信する仕組み
	for i := 0; i < 2; i++ {
		// http.Get() の代わりに、作成したクライアントのGetメソッドを使用
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
