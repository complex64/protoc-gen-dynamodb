package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

type Context struct {
	flags flag.FlagSet
	files map[*protogen.File]any
}

type out struct {
	f *protogen.GeneratedFile
}

func NewContext(flags flag.FlagSet) *Context {
	return &Context{
		flags: flags,
		files: make(map[*protogen.File]any),
	}
}

func (c *Context) Add(f *protogen.File) {
	c.files[f] = struct{}{}
}

func (c *Context) Generate(plugin *protogen.Plugin) error {
	if len(c.files) == 0 {
		return nil
	}
	for f := range c.files {
		c.generateFile(plugin, f)
	}
	return nil
}

func (c *Context) generateFile(plugin *protogen.Plugin, input *protogen.File) {
	const extension = "_dynamodb.pb.go"
	var (
		filename = input.GeneratedFilenamePrefix + extension
		file     = plugin.NewGeneratedFile(filename, input.GoImportPath)
	)
	fx := &fileContext{
		flags:  c.flags,
		plugin: plugin,
		in:     input,
		out:    file,
	}
	fx.generate()
}
