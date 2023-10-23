package serialization

import "encoding/json"

type JSONSerializer struct {
}

func (s *JSONSerializer) Unmarshal(data []byte, message interface{}) error {
	return json.Unmarshal(data, message)
}

func (s *JSONSerializer) Marshal(message interface{}) ([]byte, error) {
	return json.Marshal(message)
}
