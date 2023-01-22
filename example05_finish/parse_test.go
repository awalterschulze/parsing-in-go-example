package example05_finish

import (
	"reflect"
	"testing"

	. "github.com/awalterschulze/parsing-in-go-example/example05_finish/ast"
	"github.com/awalterschulze/parsing-in-go-example/example05_finish/gen_parser/lexer"
	"github.com/awalterschulze/parsing-in-go-example/example05_finish/gen_parser/parser"
)

func TestPass(t *testing.T) {
	input := []byte("digraph Parsing { File -> Scanner [label = Tokens]	}")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	got, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
	}
	want := &Graph{Name: "Parsing", Edges: []*Edge{{Src: "File", Dst: "Scanner", Attrs: Attrs{"label": "Tokens"}}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected ast to be %#v, but got %#v", want, got)
	}
}

func TestBig(t *testing.T) {
	input := []byte(`
	digraph G {
    File -> Scanner
    Scanner -> Parser [label = Tokens]
    Parser -> AST
    rankdir=LR
    File [shape=note]
    Parser [
        shape = record,
        style = rounded
    ]
    AST [shape=diamond]
}
	`)
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	_, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
	}
}
