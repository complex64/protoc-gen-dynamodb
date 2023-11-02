package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	extension = "_dynamodb.pb.go"
)

func New(flags flag.FlagSet) *Generator {
	return &Generator{
		flags: flags,
	}
}

type Generator struct {
	flags flag.FlagSet
}

func (g *Generator) Generate(p *protogen.Plugin) error {
	ctx := NewContext(g.flags)
	for _, f := range p.Files {
		if !f.Generate {
			continue
		}
		ctx.Add(f)
	}
	return ctx.Generate(p)
}
