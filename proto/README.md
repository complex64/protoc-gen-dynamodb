# About

Shows the options you can use in your `.proto` files.

# Usage

Import `options.proto` from your `.proto` files and set options to control what `protoc-gen-dynamodb` generates.

```proto
import "ddb/options.proto";
```

Your editor and build setup needs to reference this file.

You have multiple options, depending what tools you use:

- Copy the file into your project, e.g. a `vendor` directory
- Use [Buf's Schema Registry (BSR)](https://docs.buf.build/bsr/introduction)

## Companion Go Module

Importing `options.proto` implies your generated Go code imports the `github.com/complex64/protoc-gen-dynamodb/ddbpb`
package.

This package is distributed as [a small Go module](ddb) with minimal dependencies.
