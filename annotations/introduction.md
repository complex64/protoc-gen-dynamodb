# Introduction

Annotations are special tags that can be added to your Protocol Buffer messages
to provide additional information about the structure and purpose of the
message.

To generate DynamoDB bindings for any one of your messages, you must set
the [document annotation](https://protoc-gen-dynamodb.complex64.dev/annotations/message#document)
on the message.

To use annotations, you will need to include the
[options proto file](https://github.com/complex64/protoc-gen-dynamodb/blob/main/proto/dynamodb/options.proto)
in your project. This file contains the definitions for all the available
annotations that can be used with `protoc-gen-dynamodb`, including the document
annotation.

The generated code imports the
[dynamodbpb package](https://github.com/complex64/protoc-gen-dynamodb/tree/main/dynamodbpb).
This package only contains the Go representation of the annotations, and does
not pull in any other dependencies. 
