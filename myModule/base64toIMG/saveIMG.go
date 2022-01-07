package base64toimg

import (
	"encoding/base64"
	"fmt"
	"myModule/converter"
	"os"
	"strings"
)

func ReqJsonToImg(request *converter.Request, filepath string) []byte {
	imgAndTag := request.Text                 // Base64になってる画像部分のみを取り出す
	base64img := splitBody(imgAndTag, ",")[1] // 頭に　data:image/png;base64,　という余計な部分がくっついてるので取り除く
	fmt.Println("DeocdeToBytesData:", base64img)

	data, _ := base64.StdEncoding.DecodeString(base64img)
	file, _ := os.Create(filepath) // encode_and_decode.jpgという名称のファイルを作成
	defer file.Close()
	file.Write(data) //encode_and_decode.jpgに対して、画像のデータを書き込み
	return data
}

func splitBody(strbody string, splits string) []string {
	return strings.Split(strbody, splits)
}
