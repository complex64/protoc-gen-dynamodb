package gen

func (mx *msgCtx) genScan() {
	//mx.genScanMethod()
	//mx.genScanOperationType()
}

func (mx *msgCtx) genScanMethod() {
	var (
		p = mx.out.P
	)
	p("func (tbl ", mx.typeNameTableStruct(), ") Scan() ", mx.typenameScanOperation(), "{")
	{
		p("return ", mx.typenameScanOperation(), "{}")
	}
	p("}")
	p()
}

func (mx *msgCtx) genScanOperationType() {
	var (
		p = mx.out.P
	)
	p("type ", mx.typenameScanOperation(), " struct {")
	p("}")
	p()
}

func (mx *msgCtx) typenameScanOperation() string {
	return mx.typeNameTableStruct() + "ScanOperation"
}
