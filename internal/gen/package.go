package gen

import (
	"flag"
	"path"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	typenameDynamoContext = "DynamoContext"
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

	// Reuse file context for the header.
	fx := &fileContext{
		flags:  px.flags,
		plugin: px.plugin,
		in:     px.sample,
		out:    file,
	}
	fx.writeCommentHeader()
	fx.writePackage()

	px.out = file
	px.genDynamoContext()
}

// genDynamoContext generates the main Dynamo() function.
func (px *packageContext) genDynamoContext() {
	var (
		p   = px.out.P
		ddb = importIdentDynamoDBAPI(px.out)
	)
	p("func Dynamo("+
		"ddb ", ddb, ","+
		") ", typenameDynamoContext, " {")
	p("return ", typenameDynamoContext, "{ddb: ddb}")
	p("}")
	p()

	p("type ", typenameDynamoContext, " struct {")
	p("ddb ", ddb)
	p("}")
	p()
}
