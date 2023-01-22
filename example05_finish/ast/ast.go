package ast

import (
	"fmt"

	"github.com/awalterschulze/parsing-in-go-example/example05_finish/gen_parser/token"
)

type Graph struct {
	Name  string
	Edges []*Edge
	Nodes []*Node
	Attrs Attrs
}

func NewGraph(s interface{}) (*Graph, error) {
	return AppendStmt(&Graph{}, s)
}

func SetGraphName(g interface{}, name interface{}) (*Graph, error) {
	g.(*Graph).Name = string(name.(*token.Token).Lit)
	return g.(*Graph), nil
}

func AppendStmt(g *Graph, s interface{}) (*Graph, error) {
	switch s.(type) {
	case *Edge:
		return &Graph{Edges: append(g.Edges, s.(*Edge))}, nil
	case *Node:
		return &Graph{Nodes: append(g.Nodes, s.(*Node))}, nil
	case Attrs:
		return &Graph{Attrs: UnionMaps(g.Attrs, s.(Attrs))}, nil
	}
	return nil, fmt.Errorf("unknown type %T", s)
}

type Edge struct {
	Src   string
	Dst   string
	Attrs Attrs
}

type Node struct {
	Name  string
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
