package main

import (
	"context"
	"fmt"
	"myModule/converter"

	//myModule(モジュール名)のmyPackage(パッケージ名)を使用する
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) error {
	request, convertErr := converter.NewFileUploaderImpl().Exec(apiRequest.Body)
	fmt.Println(apiRequest)
	if convertErr != nil {
		return convertErr
	}
	fmt.Println(request)
	return nil
}

//func executeFunction() {
//	myPackage.SayHello() //
//}
//
//type MyEvent struct {
//	Name string `json:name`
//	Age  int
//}
//
//func HandleRequest(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (string, error) {
//	fmt.Println("event:", apiRequest)
//	return fmt.Sprintf("Hello %s!", "test"), nil
//}

func main() {
	lambda.Start(Handler)
	//lambda.Start(executeFunction)
	//executeFunction()
}
