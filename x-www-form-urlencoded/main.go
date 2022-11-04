package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// curl -d test=value http://localhost:18888/
func main() {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
