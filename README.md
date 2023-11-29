# Welcome

The `protoc-gen-dynamodb` [Protocol Buffer](https://developers.google.com/protocol-buffers)
("proto") compiler plugin generates a convenient Go API to map proto messages to
DynamoDB items and vice versa.

The plugin works well together with [buf](https://docs.buf.build/introduction)
or the stock [protoc compiler](https://grpc.io/docs/protoc-installation/).

Express your DynamoDB document structure as Protocol Buffer _messages_, annotate
where necessary, and then generate a Go API without having to write boilerplate
code.

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

Checkout the [examples](https://protoc-gen-dynamodb.complex64.dev/examples) for
code samples that demonstrate some of the plugin's features and use cases.

## About

[Joe](https://github.com/complex64) created this plugin in 2023 to support a few
side projects that needed simple, fast, and low-cost data persistence.

Contributions from the community are most welcome.
See [how to contribute](https://github.com/complex64/protoc-gen-dynamodb/blob/main/CONTRIBUTING.md).

If you happen to find this project useful and would like to support its
development, [show your appreciation through a small donation](https://github.com/sponsors/complex64).
Thank you!
