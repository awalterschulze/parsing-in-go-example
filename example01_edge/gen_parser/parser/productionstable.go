// Code generated by gocc; DO NOT EDIT.

package parser

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
		String: `S' : Grammar	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Grammar : id "->" id	<<  >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
}