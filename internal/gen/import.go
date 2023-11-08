package gen

import "google.golang.org/protobuf/compiler/protogen"

var (
	importContext          = newImport("context", "Context")
	importIdentDynamoDBAPI = newImport("github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface", "DynamoDBAPI")
)

func newImport(path protogen.GoImportPath, name string) func(f *protogen.GeneratedFile) string {
	return func(f *protogen.GeneratedFile) string {
		return f.QualifiedGoIdent(protogen.GoIdent{
			GoName:       name,
			GoImportPath: path,
		})
	}
}
