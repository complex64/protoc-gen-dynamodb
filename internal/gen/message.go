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
	mx.genInterfaceType()
	mx.genStructType()
	mx.genStructMethods()
}

func (mx *msgCtx) genInterfaceType() {
	var (
		p   = mx.out.P
		ctx = importIdentContext(mx.out)
	)
	p("type ", mx.typeNameWrapperInterface(), " interface {")
	{
		for _, field := range mx.Fields {
			mx.genInterfaceGetter(field)
		}
		for _, field := range mx.Fields {
			mx.genInterfaceSetter(field)
		}

		p("Reload(ctx ", ctx, ") error")
	}

	p("}")
	p()
}

func (mx *msgCtx) genInterfaceGetter(field *protogen.Field) {
	var (
		p = mx.out.P
	)
	p(field.GoName, "() ", field.Desc.Kind().String())
}

func (mx *msgCtx) genInterfaceSetter(field *protogen.Field) {
	var (
		p = mx.out.P
	)
	p("Set", field.GoName, "(value ", field.Desc.Kind().String(), ")")
}

func (mx *msgCtx) genStructType() {
	var (
		p = mx.out.P
	)
	p("type ", mx.typeNameWrapperStruct(), " struct {")
	{
		p("proto *", mx.GoIdent.GoName)

		p("changed bool")
		for _, field := range mx.Fields {
			p("set", field.GoName, " bool")
		}
	}
	p("}")
	p()
}

func (mx *msgCtx) genStructMethods() {
	for _, field := range mx.Fields {
		mx.genStructGetter(field)
	}
	for _, field := range mx.Fields {
		mx.genStructSetter(field)
	}
	mx.genStructMethodReload()
}

func (mx *msgCtx) genStructGetter(field *protogen.Field) {
	var (
		p = mx.out.P
	)
	p("func (x *", mx.typeNameWrapperStruct(), ") ", field.GoName, "() ", field.Desc.Kind().String(), " {")
	p("return ", "x.proto.Get", field.GoName, "()")
	p("}")
	p()
}

func (mx *msgCtx) genStructSetter(field *protogen.Field) {
	var (
		p = mx.out.P
	)
	p("func (x *", mx.typeNameWrapperStruct(), ") Set", field.GoName, "(value ", field.Desc.Kind().String(), ") {")
	p("x.changed = true")
	p("x.set", field.GoName, " = true")
	p("x.proto.", field.GoName, " = value")
	p("}")
	p()
}

func (mx *msgCtx) genStructMethodReload() {
	var (
		p   = mx.out.P
		ctx = importIdentContext(mx.out)
	)
	p("func (x *", mx.typeNameWrapperStruct(), ") Reload(ctx ", ctx, ") error {")
	{
		p("return nil")
	}
	p("}")
	p()
}

func (mx *msgCtx) typeNameWrapperInterface() string {
	return "DynamoDB" + mx.GoIdent.GoName
}

func (mx *msgCtx) typeNameWrapperStruct() string {
	return "dynamoDB" + mx.GoIdent.GoName
}

func getMessageOptions(message *protogen.Message) *dynampdbpb.MessageOptions {
	opts := message.Desc.Options()
	ext, ok := proto.GetExtension(opts, dynampdbpb.E_Message).(*dynampdbpb.MessageOptions)
	if ok && ext != nil {
		return ext
	}
	return &dynampdbpb.MessageOptions{}
}
