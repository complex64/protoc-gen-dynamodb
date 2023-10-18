package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

type Context struct {
	flags flag.FlagSet
	files []*protogen.File
}

func (c Context) Add(f *protogen.File) {
	c.files = append(c.files, f)
}

func (c Context) Generate(flags flag.FlagSet) error {
	c.flags = flags
	return c.generate()
}

func (c Context) generate() error {
	return nil
}
