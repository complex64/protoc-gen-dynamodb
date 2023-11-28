# protoc-gen-dynamodb

[![Tests](https://github.com/complex64/protoc-gen-dynamodb/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-dynamodb/actions/workflows/tests.yml)
[![Linters](https://github.com/complex64/protoc-gen-dynamodb/actions/workflows/linters.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-dynamodb/actions/workflows/linters.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/complex64/protoc-gen-dynamodb)](https://goreportcard.com/report/github.com/complex64/protoc-gen-dynamodb)
[![Maintainability](https://api.codeclimate.com/v1/badges/6f0c7fab9bf010198e22/maintainability)](https://codeclimate.com/github/complex64/protoc-gen-dynamodb/maintainability)
[![Go Reference](https://pkg.go.dev/badge/github.com/complex64/protoc-gen-dynamodb.svg)](https://pkg.go.dev/github.com/complex64/protoc-gen-dynamodb)

The `protoc-gen-dynamodb`  [Protocol Buffer](https://developers.google.com/protocol-buffers)
("proto") _compiler plugin_ generates a convenient Go API to map proto messages to DynamoDB items and vice versa.

The plugin works well together with [buf](https://docs.buf.build/introduction) or the
stock [protoc compiler](https://grpc.io/docs/protoc-installation/).

## Documentation

- [Installation](https://complex64.gitbook.io/protoc-gen-go-dynamodb/installation)
