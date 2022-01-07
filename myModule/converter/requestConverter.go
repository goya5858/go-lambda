package converter

/* 受け取ったRequest.Bodyを
 type Request struct {
	Line_count string // 特に意味のないデータ　複数を扱うための仮のもの
	Text       string //base64でエンコードされたデータ
}
形式に展開して返すコンバーターの設定
*/

import (
	"encoding/json"
	"fmt"
)

// Request
// フィールドは大文字から始める
type Request struct {
	Line_count string // 特に意味のないデータ　複数を扱うための仮のもの
	Text       string //base64でエンコードされたデータ
}

type requestConverterImpl struct{}

type requestConverter interface {
	Exec(requestBody string) (*Request, error)
}

func (impl requestConverterImpl) Exec(requestBody string) (*Request, error) {
	jsonBytes := []byte(requestBody)
	req := new(Request)
	fmt.Println("jsonBytes:", jsonBytes)
	fmt.Println("Before req:", req)

	if unmarshalErr := json.Unmarshal(jsonBytes, &req); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	fmt.Println("After req:", req)
	return req, nil
}

func NewFileUploaderImpl() requestConverter {
	return requestConverterImpl{}
}
