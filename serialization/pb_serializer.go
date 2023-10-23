package serialization

import (
	"errors"
	"google.golang.org/protobuf/proto"
)

type PBSerializer struct{}

func (s *PBSerializer) Unmarshal(data []byte, message interface{}) error {
	if message == nil {
		return nil
	}

	msg, ok := message.(proto.Message)
	if !ok {
		return errors.New("unmarshal fail: body not protobuf message")
	}

	return proto.Unmarshal(data, msg)
}

func (s *PBSerializer) Marshal(message interface{}) ([]byte, error) {
	msg, ok := message.(proto.Message)
	if !ok {
		return nil, errors.New("marshal fail: body not protobuf message")
	}

	return proto.Marshal(msg)
}
