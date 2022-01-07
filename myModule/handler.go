package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"myModule/base64toIMG"
	"myModule/converter"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//myModule(モジュール名)のmyPackage(パッケージ名)を使用する
)

func Handler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var header map[string]string = map[string]string{
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
		"Content-Type":                 "image/*",
	}

	fmt.Println("allBody", apiRequest.Body)
	request, convertErr := converter.NewFileUploaderImpl().Exec(apiRequest.Body)
	if convertErr != nil {
		fmt.Println(convertErr)
		res := events.APIGatewayProxyResponse{
			StatusCode:      500,
			Headers:         header,
			Body:            "",
			IsBase64Encoded: true,
		}
		return res, convertErr
	}

	filepath := "/tmp/encode_and_decord.jpg" // /tmp/以下に保存しないとエラー
	data := base64toIMG.ReqJsonToImg(request, filepath)
	fmt.Println("ByteFile:", data)

	file, _ := os.Open(filepath)
	fmt.Println("OpenFile")
	defer file.Close()

	fi, err := file.Stat() //FileInfo interface
	if err != nil {
		fmt.Println("file.Stat Error: ", err)
	}

	byte_data := make([]byte, fi.Size())
	file.Read(byte_data)
	fmt.Println("ReadFile")
	base64img := base64.StdEncoding.EncodeToString(byte_data)
	res := events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         header,
		Body:            base64img,
		IsBase64Encoded: true,
	}
	return res, nil
}

func splitBody(strbody string, splits string) []string {
	return strings.Split(strbody, splits)
}

func main() {
	lambda.Start(Handler)
	//lambda.Start(executeFunction)
	//executeFunction()
}

/*
apiRequest.Body
{
    "line_count": null,
    "text": "data:image/png;base64,<Base64でエンコードされた画像>"
}
の形式で渡される
*/
