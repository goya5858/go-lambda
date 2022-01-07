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

/*
apiRequest.Body
{
    "line_count": null,
    "text": "data:image/png;base64,<Base64でエンコードされた画像>
}
の形式で渡される
*/
func Handler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) error {
	fmt.Println(apiRequest)
	fmt.Println("allBody", apiRequest.Body)
	arr := splitBody(apiRequest.Body)
	fmt.Println("content Body:", arr[2])

	request, convertErr := converter.NewFileUploaderImpl().Exec(arr[2])
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
