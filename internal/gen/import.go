package gen

import "google.golang.org/protobuf/compiler/protogen"

var (
	importIdentContext                = newImport("context", "Context")
	importIdentDynamoDBAPI            = newImport("github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface", "DynamoDBAPI")
	importIdentAwsGetItemInput        = newImport("github.com/aws/aws-sdk-go/service/dynamodb", "GetItemInput")
	importIdentAwsString              = newImport("github.com/aws/aws-sdk-go/aws", "String")
	importIdentAwsBool                = newImport("github.com/aws/aws-sdk-go/aws", "Bool")
	importIdentDynamoDBAttributeValue = newImport("github.com/aws/aws-sdk-go/service/dynamodb", "AttributeValue")
)

func newImport(path protogen.GoImportPath, name string) func(f *protogen.GeneratedFile) string {
	return func(f *protogen.GeneratedFile) string {
		return f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       name,
			GoImportPath: path,
		})
	}
}
