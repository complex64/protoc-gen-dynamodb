# Welcome

The `protoc-gen-dynamodb` Protocol Buffer _compiler plugin_ generates convenient Go bindings for DynamoDB from your `.proto` files.

You describe your DynamoDB items as Protocol Buffer _messages_,
annotate where necessary, and then generate a Go API without having to write boilerplate code.

Go from this:

```protobuf
import "dynamodb/options.proto";
message Book {
    string isbn = 1  [(dynamodb.field).partition_key = true];
    string title = 2;
    string author = 3;
}
```

To this:

```go
// Import generated package:
// import appv1 "github.com/myorg/myapp/pkg/app/v1
books := appv1.DynamoDB(ddb).BookTable("novels")
book, _ := books.GetItem("978-0547928227").Execute(ctx)
fmt.Println(book.Title())
// The Hobbit
```



Head over to the [installation guide](installation.md) to get started.
