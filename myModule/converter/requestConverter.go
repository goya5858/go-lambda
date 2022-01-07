package converter

import "encoding/json"

// Request
type Request struct {
	line_count string `json:"image"`   // 特に意味のないデータ　複数を扱うための仮のもの
	text       int    `json:"user_id"` //base64でエンコードされたデータ
}

type requestConverterImpl struct{}

type requestConverter interface {
	Exec(requestBody string) (*Request, error)
}

func (impl requestConverterImpl) Exec(requestBody string) (*Request, error) {
	jsonBytes := []byte(requestBody)
	req := new(Request)

	if unmarshalErr := json.Unmarshal(jsonBytes, req); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return req, nil
}

func NewFileUploaderImpl() requestConverter {
	return requestConverterImpl{}
}
