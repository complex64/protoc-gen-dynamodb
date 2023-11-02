package gen

import "google.golang.org/protobuf/compiler/protogen"

type msgCtx struct {
	// file the message belongs to
	fileContext
	// the message itself
	*protogen.Message
}

func (mx *msgCtx) gen() {
	mx.genTableMethodsAndTypes()
}

func (mx *msgCtx) genTableMethodsAndTypes() {
	mx.genTableType()
	mx.genTableMethod()
}

func (mx *msgCtx) genTableMethod() {
	var (
		p = mx.out.P
	)
	p("func (ctx ", typenameDynamoContext, ") ", mx.GoIdent.GoName+"Table ("+
		"name string,",
		") ", mx.typenameTableStruct(), " {")
	{
		p("return ", mx.typenameTableStruct(), "{")
		{
			p("ddb: ctx.ddb,")
			p("name: name,")
		}
		p("}")
	}
	p("}")
	p()
}

func (mx *msgCtx) genTableType() {
	mx.genTableStruct()
	mx.genTableMethods()
}

func (mx *msgCtx) genTableStruct() {
	var (
		p = mx.out.P
	)
	p("type ", mx.typenameTableStruct(), " struct {")
	{
		p("ddb ", importIdentDynamoDBAPI(mx.out))
		p("name string")
	}
	p("}")
	p()
}

func (mx *msgCtx) genTableMethods() {
	mx.genTableScan()
}

func (mx *msgCtx) genTableScan() {
	mx.genTableScanMethod()
	mx.genTableScanOperationType()
}

func (mx *msgCtx) genTableScanMethod() {
	var (
		p = mx.out.P
	)
	p("func (tbl ", mx.typenameTableStruct(), ") Scan() ", mx.typenameTableScanOperation(), "{")
	{
		p("return ", mx.typenameTableScanOperation(), "{}")
	}
	p("}")
	p()
}

func (mx *msgCtx) genTableScanOperationType() {
	var (
		p = mx.out.P
	)
	p("type ", mx.typenameTableScanOperation(), " struct {")
	p("}")
	p()
}

func (mx *msgCtx) typenameTableStruct() string {
	return "Dynamo" + mx.GoIdent.GoName + "Table"
}

func (mx *msgCtx) typenameTableScanOperation() string {
	return mx.typenameTableStruct() + "ScanOperation"
}
