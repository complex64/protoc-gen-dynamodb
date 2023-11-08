package gen

func (mx *msgCtx) genUpdate() {
	//mx.genUpdateMethod()
	//mx.genUpdateOperationType()
}

func (mx *msgCtx) genUpdateMethod() {
	var (
		p = mx.out.P
	)
	p("func (tbl ", mx.typeNameTableStruct(), ") Update() ", mx.typenameUpdateOperation(), "{")
	{
		p("return ", mx.typenameUpdateOperation(), "{}")
	}
	p("}")
	p()
}

func (mx *msgCtx) genUpdateOperationType() {
	var (
		p = mx.out.P
	)
	p("type ", mx.typenameUpdateOperation(), " struct {")
	p("}")
	p()
}

func (mx *msgCtx) typenameUpdateOperation() string {
	return mx.typeNameTableStruct() + "UpdateOperation"
}
