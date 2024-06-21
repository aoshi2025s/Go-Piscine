## Go-Piscine
printrune.goや一部のmain.go内のコードは自分で書いてないもの(コピペしたもの)が含まれています。

## Golangの実行コンテナへ入る
実行すると、ホストPCのカレントディレクトリ以下のファイルが全てコンテナへマウントされます。
```
./boot-go-container.sh
```

## 実行
```bash
cd goXX/exXX/
```
```bash
go mod init exXX
go run .
```
