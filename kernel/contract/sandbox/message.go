package sandbox

import (
	"bytes"
	"reflect"

	"github.com/golang/protobuf/proto"
)

var (
	protoIface = reflect.TypeOf((*proto.Message)(nil)).Elem()
)

func isMsgEqual(reqHead, reqIncome proto.Message) bool {
	encodeHead, err := encodeMsg(reqHead)
	if err != nil {
		return false
	}
	encodeIncome, err := encodeMsg(reqIncome)
	if err != nil {
		return false
	}
	return bytes.Equal(encodeHead, encodeIncome)
}

func encodeMsg(req proto.Message) ([]byte, error) {
	var buf proto.Buffer
	buf.SetDeterministic(true)
	err := buf.EncodeMessage(req)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
