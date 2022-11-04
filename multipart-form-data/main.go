package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// curl -F "name=Michel" -F "thumbnail=@photo.jpg" http://localhost:18888
// テキストデータとファイルの2つのデータを送信する
// 以下の方法でもファイルを送信することはできるが、各ファイルのContent-Typeはvoid型とも言える application/octet-stream型になる
// Content-TypeやMIMEタイプを指定する場合には、 /multipart-form-data/with-MINE を参照
func main() {
	var buffer bytes.Buffer                                            // マルチパート部を組み立てた後のバイト列を格納するバッファを宣言
	writer := multipart.NewWriter(&buffer)                             // マルチパートを組み立てるwriterを作成 値ではなく、ポインタを渡す
	writer.WriteField("name", "Michel")                                // ファイル以外のフィールドは、WriteFieldを使って登録できる
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg") // ファイル書き込みのio.Writerを作成
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg") // ファイルを開く
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile) // 開いたファイルの全コンテンツを、ファイル書き込み用のio.Writerにコピーする
	writer.Close()                // Closeすることで、バッファにすべてを書き込む

	// Content-Type には バウンダリー文字列を入れる必要がある。
	// writer.FormDataContentType() または "multipart/form-data; boundary=" + writer.Boundary() で生成できる
	// バウンダリー文字列はmultipart.Writerオブジェクトが内部で乱数を使って生成する。
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
