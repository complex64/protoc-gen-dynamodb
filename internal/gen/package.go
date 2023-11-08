package gen

import (
	"flag"
	"path"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	typeNameDynamoDBContext = "DynamoDBContext"
)

type packageContext struct {
	// flags from the command line
	flags flag.FlagSet

	// protoc plugin handle
	plugin *protogen.Plugin

	// an input proto file from this package
	sample *protogen.File

	// output file we generate
	out *protogen.GeneratedFile
}

func (px *packageContext) gen() {
	var (
		dir      = path.Dir(px.sample.GeneratedFilenamePrefix)
		filename = path.Join(dir, "dynamodb.pb.go")
		file     = px.plugin.NewGeneratedFile(filename, px.sample.GoImportPath)
	)

	fx := &fileContext{
		flags:  px.flags,
		plugin: px.plugin,
		in:     px.sample,
		out:    file,
	}
	fx.writeCommentHeader()
	fx.writePackage()

	px.out = file
	px.genWithDynamoDBFunc()
	px.genReturnConsumedCapacityType()
}

func (px *packageContext) genWithDynamoDBFunc() {
	var (
		p          = px.out.P
		clientType = importIdentDynamoDBAPI(px.out)
	)
	p("func WithDynamoDB(client ", clientType, ") ", typeNameDynamoDBContext, " {")
	p("return ", typeNameDynamoDBContext, "{client: client}")
	p("}")
	p()

	p("type ", typeNameDynamoDBContext, " struct {")
	p("client ", clientType)
	p("}")
	p()
}

func (px *packageContext) genReturnConsumedCapacityType() {
	var (
		p    = px.out.P
		name = "ReturnConsumedCapacityLevel"

		ddbPkg  protogen.GoImportPath = "github.com/aws/aws-sdk-go/service/dynamodb"
		indexes                       = newImport(ddbPkg, "ReturnConsumedCapacityIndexes")(px.out)
		total                         = newImport(ddbPkg, "ReturnConsumedCapacityTotal")(px.out)
		none                          = newImport(ddbPkg, "ReturnConsumedCapacityNone")(px.out)
	)
	p("type ", name, " struct { string }")
	p()

	p("var (")
	p("ReturnIndexes = ", name, "{", indexes, "}")
	p("ReturnTotal = ", name, "{", total, "}")
	p("ReturnNone = ", name, "{", none, "}")
	p(")")
	p()
}
