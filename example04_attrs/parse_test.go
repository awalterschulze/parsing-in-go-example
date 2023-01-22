package example04_attrs

import (
	"reflect"
	"testing"

	. "github.com/awalterschulze/parsing-in-go-example/example04_attrs/ast"
	"github.com/awalterschulze/parsing-in-go-example/example04_attrs/gen_parser/lexer"
	"github.com/awalterschulze/parsing-in-go-example/example04_attrs/gen_parser/parser"
)

func TestPass(t *testing.T) {
	input := []byte("digraph { File -> Scanner [label = Tokens]	}")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	got, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
	}
	want := &Graph{Edges: []*Edge{{Src: "File", Dst: "Scanner", Attrs: Attrs{"label": "Tokens"}}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected ast to be %#v, but got %#v", want, got)
	}
}
