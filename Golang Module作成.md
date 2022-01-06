・モジュールディレクトリに移動
```
>> cd ./TestLambda
```

・モジュールの作成
```
>> go mod init <任意のモジュール名>
```
で任意の名称のモジュールを作成できる  


・パッケージの作成
go.modが存在するディレクトリ以下で、任意のGoファイルのパッケージ名を
```go
// myModule/myPackage/hello.go
package myPackage

import "fmt"

func SayHello() {
	fmt.Println("Hello Lambda")
}
```
のようにPackage名を付ける  
<span style="color: red; ">
パッケージ名は ファイルの配置されるディレクトリと同じが望ましい
</span>

・メインのファイルで使用するとき、
```go
// myModule/handler.go
package main

import (
	"myModule/myPackage" //自作のmyModule(モジュール名)の、myPackage(パッケージ名)を使用するということ
	// <packageName> "<ModuleName>/<使用するpackageのディレクトリ>" 
	// packageの置いてあるディレクトリ名とpackage名が異なる場合

	"github.com/aws/aws-lambda-go/lambda" //外部モジュール
)
```
のように使用する