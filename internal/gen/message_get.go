package gen

func (mx *msgCtx) genGet() {
	mx.genGetMethod()
	mx.genGetOpType()
	mx.genGetOpMethods()
}

func (mx *msgCtx) genGetMethod() {
	var (
		p      = mx.out.P
		pkName = mx.paramNamePartitionKey()
		skName = mx.paramNameSortKey()
	)

	// todo: support types keys; string, number, binary, bool ONLY, for now.

	args := pkName + " string"
	if mx.sortKey != nil {
		args += ", " + string(skName) + " string"
	}

	p("func (tx ", mx.typeNameTableStruct(), ") GetItem(", args, ") ", mx.typeGetOp(), "{")
	{
		p("return ", mx.typeGetOp(), "{")
		p("tableContext: tx,")
		p("partitionKey: ", pkName, ",")
		if mx.sortKey != nil {
			p("sortKey: ", skName, ",")
		}
		p("}")
	}
	p("}")
	p()
}

func (mx *msgCtx) paramNamePartitionKey() string {
	if mx.partitionKey == nil {
		panic("BUG: partition key is nil")
	}
	n := mx.partitionKey.GoName
	return string(n[0]+32) + n[1:]
}

func (mx *msgCtx) paramNameSortKey() string {
	if mx.sortKey == nil {
		return ""
	}
	n := mx.sortKey.GoName
	return string(n[0]+32) + n[1:]
}

func (mx *msgCtx) genGetOpType() {
	// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html
	var (
		p = mx.out.P
	)
	p("type ", mx.typeGetOp(), " struct {")
	p("tableContext ", mx.typeNameTableStruct())
	p("partitionKey string")
	p("sortKey string")
	p("consistentRead bool")
	p("returnConsumedCapacity string")
	p("}")
	p()
}

func (mx *msgCtx) genGetOpMethods() {
	mx.genGetOpSetConsistentRead()
	mx.genGetOpReturnConsumedCapacity()
	mx.genGetOpExecute()
}

func (mx *msgCtx) genGetOpSetConsistentRead() {
	var (
		p = mx.out.P
	)
	p("// SetConsistentRead configures the GetItem operation to use strongly consistent reads.")
	p("// Otherwise, the operation uses eventually consistent reads.")
	p("func (op ", mx.typeGetOp(), ") SetConsistentRead() (", mx.typeGetOp(), "){")
	p("op.consistentRead = true")
	p("return op")
	p("}")
	p()
}

func (mx *msgCtx) genGetOpReturnConsumedCapacity() {
	var (
		p = mx.out.P
	)
	p("func (op ", mx.typeGetOp(), ") ReturnConsumedCapacity(level ReturnConsumedCapacityLevel) (", mx.typeGetOp(), "){")
	p("op.returnConsumedCapacity = string(level)")
	p("return op")
	p("}")
	p()
}

func (mx *msgCtx) genGetOpExecute() {
	var (
		p            = mx.out.P
		ctx          = importIdentContext(mx.out)
		getItemInput = importIdentAwsGetItemInput(mx.out)
		awsStr       = importIdentAwsString(mx.out)
		awsBool      = importIdentAwsBool(mx.out)
		ddbAttrVal   = importIdentDynamoDBAttributeValue(mx.out)
	)

	p("func (op ", mx.typeGetOp(), ") Execute("+
		"ctx ", ctx, ","+
		") (",
		mx.typeNameWrapperInterface(), ",",
		"error,",
		") {")
	{
		p("input := &", getItemInput, "{")
		{
			p("TableName: ", awsStr, "(op.tableContext.name),")
			p("Key: map[string]*", ddbAttrVal, "{")
			{

				p("\"", mx.partitionKey.Desc.Name(), "\": {")
				p("S: ", awsStr, "(op.partitionKey),")
				p("},")

				if mx.sortKey != nil {
					p("\"", mx.sortKey.Desc.Name(), "\": {")
					p("S: ", awsStr, "(op.sortKey),")
					p("},")

				}
			}
			p("},")
			p("ConsistentRead: ", awsBool, "(op.consistentRead),")
			p("ReturnConsumedCapacity: ", awsStr, "(op.returnConsumedCapacity),")
		}
		p("}")

		p("output, err := op.tableContext.client.GetItemWithContext(ctx, input)")
		p("if err != nil {")
		p("return nil, err")
		p("}")

		p("proto, err := ", mx.funcNameUnmarshaler(), "(output.Item)")
		p("if err != nil {")
		p("return nil, err")
		p("}")
		p("return &", mx.typeNameWrapperStruct(), "{proto: proto}, nil")
	}
	p("}")
	p()
}

func (mx *msgCtx) typeGetOp() string {
	return "DynamoGet" + mx.GoIdent.GoName + "Operation"
}
