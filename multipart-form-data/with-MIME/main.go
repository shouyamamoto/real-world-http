package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michel")

	// multipart.Writer.WriteField や multipart.Writer.CreateFormFileは、マルチパートの構成要素である
	// パートそのものに触れずにコンテンツを作成できる高度なAPI。各パートの持つヘッダーも自動設定された。
	// 以下のコードは上記それぞれの中で行っている詳細な処理を取り出すことで、任意のContent-Typeを指定できるようにしている。
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
