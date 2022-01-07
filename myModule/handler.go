package main

import (
	"context"
	"fmt"
	"myModule/converter"
	"strings"

	//myModule(モジュール名)のmyPackage(パッケージ名)を使用する
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) error {
	request, convertErr := converter.NewFileUploaderImpl().Exec(apiRequest.Body)
	fmt.Println(apiRequest)
	fmt.Println("allBody", apiRequest.Body)
	arr := splitBody(apiRequest.Body)
	fmt.Println("content Body:", arr[1])
	if convertErr != nil {
		return convertErr
	}
	fmt.Println(request)
	return nil
}

func splitBody(strbody string) []string {
	return strings.Split(strbody, ",")
}

func main() {
	lambda.Start(Handler)
	//lambda.Start(executeFunction)
	//executeFunction()
}
