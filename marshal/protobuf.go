package marshal

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

type ProtoMarshaler struct{}

func (*ProtoMarshaler) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, errors.New("invalid protobuf message")
	}
	return proto.Marshal(pb)
}

func (*ProtoMarshaler) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return errors.New("invalid protobuf message")
	}
	return proto.Unmarshal(data, pb)
}

func (*ProtoMarshaler) String() string {
	return "protobuf"
}

func (*ProtoMarshaler) ContentType(_ interface{}) string {
	return "application/protobuf"
}
