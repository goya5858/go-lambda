package main

import (
	"context"
	"fmt"
	"myModule/myPackage" //myModule(モジュール名)のmyPackage(パッケージ名)を使用する

	"github.com/aws/aws-lambda-go/lambda"
)

func executeFunction() {
	myPackage.SayHello() //
}

type MyEvent struct {
	Name string `json:name`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fmt.Println(event)
	return fmt.Sprintf("Hello %s!", event.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
	//lambda.Start(executeFunction)
	//executeFunction()
}
