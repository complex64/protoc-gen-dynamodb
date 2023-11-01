package proto

//go:generate go install github.com/bufbuild/buf/cmd/buf@v1.27.1
//go:generate buf lint
//go:generate find .. -name '*.pb.go' -delete
//go:generate buf generate
