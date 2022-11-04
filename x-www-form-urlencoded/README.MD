# x-www-form-urlencoded
x-www-form-urlencodedは、シンプルなフォームの送信である。

```html
<form method="POST">
    <input name="title">
    <input name="author">
    <input name="submit">
</form>
```

このようなフォームをシンプルなフォームとしている。  
以下のようなcurlコマンドを使って、上記と同じような送信ができる。  

```
$ curl --http1.0 -d title="The Art of Community" -d author="Jono Bacon" http://localhost:18888
```

curlコマンドの`-d`オプションを使ってフォームで送信するデータを設定できる。  
curlコマンドは`-d`オプションが指定されると、ブラウザと同じようにヘッダーとして`Content-Type: application/x-www-form-urlencoded`を設定する。  
  
また、`-d`オプションだと指定された文字列をそのまま送信してしまう。空白や&などを文字列として送信する場合、読み込み先で正しく元のデータセットを復元できなくなる。  
そのため、空白や&がある文字列を送信する場合は、`--data-urlencoded`を使うと良い。  
空白は`%20`に、&は`%26`に変換される。  
  
```
$ curl --http1.0 --data-urlencoded title="Head First PHP & MySQL" --data-urlencoded author="Lynn, Michael" http://localhost:18888

title=Head%20First%20PHP%20%20%26%20MySQL&author=Lynn%2c%20Michael
```
