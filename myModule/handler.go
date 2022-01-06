package main

import (
	"myModule/myPackage" //myModule(モジュール名)のmyPackage(パッケージ名)を使用する

	"github.com/aws/aws-lambda-go/lambda"
)

func executeFunction() {
	myPackage.SayHello() //
}

func main() {
	lambda.Start(executeFunction)
	//executeFunction()
}
