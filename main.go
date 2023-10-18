package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/complex64/protoc-gen-go-dynamodb/internal/gen"
	"github.com/complex64/protoc-gen-go-dynamodb/internal/zerolog"
)

func main() {
	zerolog.Init()
	handleFlags()

	var flags flag.FlagSet
	opts := protogen.Options{ParamFunc: flags.Set}
	// Run calls `flags.Set(param, value)`
	// for each `--go_dynamodb_out=<param1>=<value1>,...`.
	opts.Run(gen.New(flags).Generate)
}

func handleFlags() {
	version := flag.Bool("version", false, "Print the version and exit.")
	flag.Parse()
	if *version {
		name := filepath.Base(os.Args[0])
		log.Info().Msg(name + " version v" + gen.Version)
		os.Exit(0)
	}
}
