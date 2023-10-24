# protoc-gen-go-dynamodb

[![Tests](https://github.com/complex64/protoc-gen-go-dynamodb/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-go-dynamodb/actions/workflows/tests.yml)
[![Linters](https://github.com/complex64/protoc-gen-go-dynamodb/actions/workflows/linters.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-go-dynamodb/actions/workflows/linters.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/complex64/protoc-gen-go-dynamodb)](https://goreportcard.com/report/github.com/complex64/protoc-gen-go-dynamodb)
[![Maintainability](https://api.codeclimate.com/v1/badges/6f0c7fab9bf010198e22/maintainability)](https://codeclimate.com/github/complex64/protoc-gen-go-dynamodb/maintainability)
[![Go Reference](https://pkg.go.dev/badge/github.com/complex64/protoc-gen-go-dynamodb.svg)](https://pkg.go.dev/github.com/complex64/protoc-gen-go-dynamodb)

Generate DynamoDB bindings for Go from your .proto files.

`protoc-gen-go-dynamodb` is a plugin for [protoc](https://grpc.io/docs/protoc-installation/)
or [buf](https://docs.buf.build/introduction),
a [Protocol Buffer](https://developers.google.com/protocol-buffers) ("proto") compiler.
