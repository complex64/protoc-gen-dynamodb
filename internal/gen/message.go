package gen

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	dynampdbpb "github.com/complex64/protoc-gen-dynamodb/dynamodbpb"
)

type msgCtx struct {
	// file the message belongs to
	fileContext

	// the message itself
	*protogen.Message

	partitionKey *protogen.Field
	sortKey      *protogen.Field
}

func (mx *msgCtx) gen() {
	if !getMessageOptions(mx.Message).Document {
		return
	}

	mx.init()
	mx.genTableMethodsAndTypes()
	mx.genWrapperMethodsAndTypes()
}

func (mx *msgCtx) init() {
	mx.determinePartitionKey()
	mx.determineSortKey()
}

func (mx *msgCtx) determinePartitionKey() {
	found := false
	for _, field := range mx.Fields {
		if getFieldOptions(field).PartitionKey {
			if found {
				mx.plugin.Error(fmt.Errorf(
					"%s: message %s has more than one partition key",
					mx.fileContext.in.Desc.Path(),
					mx.GoIdent.GoName,
				))
				return
			}

			mx.partitionKey = field
			found = true
		}
	}

	if !found {
		mx.plugin.Error(fmt.Errorf(
			"%s: message %s has no partition key",
			mx.fileContext.in.Desc.Path(),
			mx.GoIdent.GoName,
		))
		return
	}

	// TODO: Ensure the partition key is of a supported type.
}

func (mx *msgCtx) determineSortKey() {
	found := false
	for _, field := range mx.Fields {
		if getFieldOptions(field).SortKey {
			if found {
				mx.plugin.Error(fmt.Errorf(
					"%s: message %s has more than one sort key",
					mx.fileContext.in.Desc.Path(),
					mx.GoIdent.GoName,
				))
				return
			}

			mx.sortKey = field
			found = true
		}
	}
}

func (mx *msgCtx) genTableMethodsAndTypes() {
	mx.genTableType()
	mx.genTableMethod()
}

func (mx *msgCtx) genTableType() {
	mx.genTableStruct()
	mx.genTableMethods()
}

func (mx *msgCtx) genTableStruct() {
	var (
		p          = mx.out.P
		clientType = importIdentDynamoDBAPI(mx.out)
	)
	p("type ", mx.typeNameTableStruct(), " struct {")
	{
		p("client ", clientType)
		p("name string")
	}
	p("}")
	p()
}

func (mx *msgCtx) genTableMethods() {
	mx.genGet()
	mx.genScan()
	mx.genUpdate()
}

func (mx *msgCtx) genTableMethod() {
	var (
		p = mx.out.P
	)
	p("func (dx ", typeNameDynamoDBContext, ") With", mx.GoIdent.GoName+"Table ("+
		"name string,",
		") ", mx.typeNameTableStruct(), " {")
	{
		p("return ", mx.typeNameTableStruct(), "{")
		{
			p("client: dx.client,")
			p("name: name,")
		}
		p("}")
	}
	p("}")
	p()
}

func (mx *msgCtx) typeNameTableStruct() string {
	return "DynamoDB" + mx.GoIdent.GoName + "TableContext"
}

func (mx *msgCtx) genWrapperMethodsAndTypes() {
	mx.genWrapperType()
	mx.genWrapperMethods()
}

func (mx *msgCtx) genWrapperType() {
	var (
		p = mx.out.P
	)
	p("type ", mx.typeNameWrapperInterface(), " interface {")

	for _, field := range mx.Fields {
		mx.genGetter(field)
	}
	for _, field := range mx.Fields {
		mx.genSetter(field)
	}

	p("}")
	p()
}

func (mx *msgCtx) genGetter(field *protogen.Field) {
	var (
		p    = mx.out.P
		name = mx.methodNameGet(field)
	)
	p(name, "() ", field.Desc.Kind().String())
}

func (mx *msgCtx) genSetter(field *protogen.Field) {
	var (
		p    = mx.out.P
		name = mx.methodNameSet(field)
	)
	p(name, "(value ", field.Desc.Kind().String(), ") ", mx.typeNameWrapperInterface())
}

func (mx *msgCtx) methodNameGet(field *protogen.Field) string {
	return getFieldGoName(field)
}

func (mx *msgCtx) methodNameSet(field *protogen.Field) string {
	return "Set" + getFieldGoName(field)
}

func (mx *msgCtx) genWrapperMethods() {
	var (
		p = mx.out.P
	)
	_ = p
}

func (mx *msgCtx) typeNameWrapperInterface() string {
	return "DynamoDB" + mx.GoIdent.GoName
}

func getMessageOptions(message *protogen.Message) *dynampdbpb.MessageOptions {
	opts := message.Desc.Options()
	ext, ok := proto.GetExtension(opts, dynampdbpb.E_Message).(*dynampdbpb.MessageOptions)
	if ok && ext != nil {
		return ext
	}
	return &dynampdbpb.MessageOptions{}
}
