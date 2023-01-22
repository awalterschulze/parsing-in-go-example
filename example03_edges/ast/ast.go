package ast

import (
	"github.com/awalterschulze/parsing-in-go-example/example03_edges/gen_parser/token"
)

type Graph struct {
	Edges []*Edge
}

func NewGraph(s interface{}) (*Graph, error) {
	return AppendStmt(&Graph{}, s)
}

func AppendStmt(g *Graph, s interface{}) (*Graph, error) {
	return &Graph{Edges: append(g.Edges, s.(*Edge))}, nil
}

type Edge struct {
	Src string
	Dst string
}

func ID(id interface{}) string {
	return string(id.(*token.Token).Lit)
}
