package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false, // 枚举是否转成整形
		EmitDefaults: true,  // 默认值
		Indent: "	", // 缩进
		OrigName: true, // proto文件中定义的原始字段名称
	}
	return marshaler.MarshalToString(message)
}
