・モジュールディレクトリに移動
```
>> cd src/TestLambda
```

・モジュールの作成
```
>> go mod init <任意のモジュール名>
```
で任意の名称のモジュールを作成できる

・パッケージの作成
go.modが存在するディレクトリ以下で、任意のGoファイルのパッケージ名を
```go
// greeting/hello.go
package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello Lambda")
}
```
のようにPackage名を付ける

・メインのファイルで使用するとき、
```go
// handler.go
package main

import (
	"TestLambda/greeting" //自作のTestLambda Moduleの、greeting Packageを使うということ

	"github.com/aws/aws-lambda-go/lambda" //外部モジュール
)
```
のように使用する