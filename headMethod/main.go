package main

import (
	"fmt"
	"net/http"
)

// curl --head https://google.com
func main() {
	resp, err := http.Head("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status) // 200 OK
	for k, v := range resp.Header {
		fmt.Printf("%v: %v\n", k, v[0])
	}
}
