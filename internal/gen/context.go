package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

type Context struct {
	flags    flag.FlagSet
	files    map[*protogen.File]any
	packages map[protogen.GoImportPath]*packageContext
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

// Generate everything.
func (c *Context) Generate(plugin *protogen.Plugin) error {
	for f := range c.files {
		c.collectPackage(plugin, f)
		c.genFile(plugin, f)
	}
	for _, p := range c.packages {
		p.gen()
	}
	return nil
}

// genFile generates a single output file from a single input .proto file.
func (c *Context) genFile(plugin *protogen.Plugin, input *protogen.File) {
	skip := true
	for _, m := range input.Messages {
		if getMessageOptions(m).Document {
			skip = false
			break
		}
	}
	if skip {
		return
	}

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
	fx.gen()
}

// collectPackage keeps track of all packages across all input files.
func (c *Context) collectPackage(plugin *protogen.Plugin, f *protogen.File) {
	if c.packages == nil {
		c.packages = make(map[protogen.GoImportPath]*packageContext)
	}
	if _, ok := c.packages[f.GoImportPath]; !ok {
		c.packages[f.GoImportPath] = &packageContext{
			flags:  c.flags,
			plugin: plugin,
			sample: f,
		}
	}
}
