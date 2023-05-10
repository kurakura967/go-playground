# Qiita

こちらの「[[Golang][並行処理] スライスへの同時アクセスによる処理速度の検証](https://qiita.com/kurakura0916/items/a45793aa66269f1d5675)」の記事を書く際に作成したコードです。

## ベンチマークテストによる検証

```bash
go test -bench . -benchmem
```

## プロファイリングによる検証

```bash
# 並行処理の場合
go run after01/main.go
go tool trace trace_concurrency01.out

# 逐次処理の場合
go run before/main.go
go tool trace trace_sequential.out 
```
