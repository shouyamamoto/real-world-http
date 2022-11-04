package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	client := http.Client{
		Transport: transport,
	}
	resp, err := client.Get("file://./main.txt")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	/*
		HTTP/1.0 200 OK
		Connection: close
		Accept-Ranges: bytes
		Content-Type: text/plain; charset=utf-8
		Last-Modified: Thu, 01 Sep 2022 13:14:20 GMT

		これはテスト用のテキストです。
	*/
}
