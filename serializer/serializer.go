package serializer

type Serializer interface {
	Unmarshal(data []byte, message interface{}) error
	Marshal(message interface{}) ([]byte, error)
}
