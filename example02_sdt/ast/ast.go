package ast

import (
	"github.com/awalterschulze/parsing-in-go-example/example02_sdt/gen_parser/token"
)

type Edge struct {
	Src string
	Dst string
}

func ID(id interface{}) string {
	return string(id.(*token.Token).Lit)
}
