package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	dynampdbpb "github.com/complex64/protoc-gen-dynamodb/dynamodbpb"
)

func getFieldOptions(field *protogen.Field) *dynampdbpb.FieldOptions {
	opts := field.Desc.Options()
	ext, ok := proto.GetExtension(opts, dynampdbpb.E_Field).(*dynampdbpb.FieldOptions)
	if ok && ext != nil {
		return ext
	}
	return &dynampdbpb.FieldOptions{}
}
