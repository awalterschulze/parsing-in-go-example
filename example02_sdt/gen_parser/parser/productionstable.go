// Code generated by gocc; DO NOT EDIT.

package parser

import . "github.com/awalterschulze/parsing-in-go-example/example02_sdt/ast"

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : Grammar	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Grammar : EdgeStmt	<<  >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `EdgeStmt : id "->" id	<< &Edge{Src: ID(X[0]), Dst: ID(X[2])}, nil >>`,
		Id:         "EdgeStmt",
		NTType:     2,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return &Edge{Src: ID(X[0]), Dst: ID(X[2])}, nil
		},
	},
}
