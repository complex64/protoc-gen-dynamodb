package gen

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
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
	var job Job
	for _, f := range p.Files {
		if !f.Generate {
			continue
		}
		job.Add(f)
	}
	return job.Generate(g.flags)
}
