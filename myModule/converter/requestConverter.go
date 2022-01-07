package converter

import (
	"encoding/json"
	"fmt"
)

// Request
type Request struct {
	line_count string `json:"line_count"` // 特に意味のないデータ　複数を扱うための仮のもの
	text       string `json:"text"`       //base64でエンコードされたデータ
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

	if unmarshalErr := json.Unmarshal(jsonBytes, req); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	fmt.Println("After req:", req)
	return req, nil
}

func NewFileUploaderImpl() requestConverter {
	return requestConverterImpl{}
}
