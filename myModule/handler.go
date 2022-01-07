package main

import (
	"context"
	"encoding/base64"
	"fmt"
	base64toimg "myModule/base64toIMG"
	"myModule/converter"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//myModule(モジュール名)のmyPackage(パッケージ名)を使用する
)

func Handler(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	header := map[string]string{} //ちゃんとMAPを初期化
	header["Access-Control-Allow-Headers"] = "Content-Type"
	header["Access-Control-Allow-Origin"] = "*"
	header["Access-Control-Allow-Methods"] = "OPTIONS,POST,GET"
	header["Content-Type"] = "image/*"

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

	filepath := "encode_and_decord.jpg"
	data := base64toimg.ReqJsonToImg(request, filepath)
	fmt.Println(data)

	file, _ := os.Open(filepath)
	defer file.Close()
	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	byte_data := make([]byte, size)
	file.Read(byte_data)
	ango := base64.StdEncoding.EncodeToString(byte_data)
	res := events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         header,
		Body:            ango,
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
