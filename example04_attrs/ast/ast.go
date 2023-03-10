package ast

import (
	"github.com/awalterschulze/parsing-in-go-example/example04_attrs/gen_parser/token"
)

type Graph struct {
	Name  string
	Edges []*Edge
}

func NewGraph(s interface{}) (*Graph, error) {
	return AppendStmt(&Graph{}, s)
}

func SetGraphName(g interface{}, name interface{}) (*Graph, error) {
	g.(*Graph).Name = string(name.(*token.Token).Lit)
	return g.(*Graph), nil
}

func AppendStmt(g *Graph, s interface{}) (*Graph, error) {
	return &Graph{Edges: append(g.Edges, s.(*Edge))}, nil
}

type Edge struct {
	Src   string
	Dst   string
	Attrs Attrs
}

func ID(id interface{}) string {
	return string(id.(*token.Token).Lit)
}

type Attrs map[string]string

func UnionMaps(maps ...Attrs) Attrs {
	union := make(Attrs)
	for _, m := range maps {
		for k, v := range m {
			union[k] = v
		}
	}
	return union
}
