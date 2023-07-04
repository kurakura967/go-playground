# deepdive into gRPC
ref: https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/intro

## セットアップ
```bash
$ brew install protobuf
$ brew install protoc-gen-go
$ go get -u google.golang.org/grpc
$ go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## protoファイルからコードを生成
```bash
cd src/api
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative hello.proto
```

## サーバーの起動とメソッドの呼び出し
```bash
$ go run cmd/server/main.go 
2023/07/03 23:26:53 start server on port: 8080

# 別のターミナルから実行
$ grpcurl -plaintext localhost:8080 list
grpc.reflection.v1alpha.ServerReflection
myapp.GreetingService

# GreetingServiceのメソッド一覧を表示
$ grpcurl -plaintext localhost:8080 list myapp.GreetingService
myapp.GreetingService.Hello

# Helloメソッドを実行
$ grpcurl -plaintext -d '{"name": "hoge"}' localhost:8080 myapp.GreetingService.Hello
{
  "message": "Hello, hoge"
}

# HelloServerStreamメソッド(サーバーストリーミング)を実行
$ grpcurl -plaintext -d '{"name": "hoge"}' localhost:8080 myapp.GreetingService.HelloServerStream
{
  "message": "[0] Hello, hoge!"
}
{
  "message": "[1] Hello, hoge!"
}
{
  "message": "[2] Hello, hoge!"
}
{
  "message": "[3] Hello, hoge!"
}
{
  "message": "[4] Hello, hoge!"
}


# HelloClientStreamメソッド(クライアントストリーミング)を実行
$ grpcurl -plaintext -d '{"name": "hoge"}{"name": "fuga"}{"name":"piyo"}' localhost:8080 myapp.GreetingService.HelloClientStream
{
  "message": "Hello, [hoge fuga piyo]!"
}

# HelloBiStreamsメソッド(双方向ストリーミング)を実行
$ grpcurl -plaintext -d '{"name": "hoge"}{"name": "fuga"}{"name":"piyo"}' localhost:8080 myapp.GreetingService.HelloBiStreams
{
  "message": "Hello, hoge"
}
{
  "message": "Hello, fuga"
}
{
  "message": "Hello, piyo"
}

```

## クライアントからの呼び出し
```bash
# サーバーの起動
$ go run cmd/server/main.go 

# クライアントの起動
$ go run cmd/client/main.go 
start gRPC Client.
1: Hello
2: HelloServerStream
3: HelloClientStream
4: HelloBidirectionalStream
5: exit
please enter >1
please enter your name
hoge
Hello, hoge
1: Hello
2: HelloServerStream
3: HelloClientStream
4: HelloBidirectionalStream
5: exit
please enter >5
bye
```

### 異なる言語(Python)のクライアントからの呼び出し
#### セットアップ
```bash
$ poetry init
$ poetrry env use 3.11.2
$ poetry add grpcio
$ poetry add grpcio-tools
```

#### protoファイルからコードを生成
```bash
poetry run python -m grpc_tools.protoc -I src/api --python_out=src/cmd/client/ --grpc_python_out=src/cmd/client/ src/api/hello.proto
```

#### Helloメソッドの呼び出し
```bash
# クライアントの起動
$ poetry run python src/cmd/client/main.py 
start gRPC Client
please enter your name
hoge
Hello, hoge
```
