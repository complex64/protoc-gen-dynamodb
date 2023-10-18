package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

type Job struct {
	flags flag.FlagSet
	files []*protogen.File
}

func (j Job) Add(f *protogen.File) {
	j.files = append(j.files, f)
}

func (j Job) Generate(flags flag.FlagSet) error {
	j.flags = flags
	return j.generate()
}

func (j Job) generate() error {
	return nil
}
