package proto

//go:generate go run github.com/bufbuild/buf/cmd/buf@v1.27.1 lint
//go:generate find .. -name '*.pb.go' -delete
//go:generate go run github.com/bufbuild/buf/cmd/buf@v1.27.1 generate
