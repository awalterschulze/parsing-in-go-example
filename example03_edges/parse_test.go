package example03_edges

import (
	"reflect"
	"testing"

	. "github.com/awalterschulze/parsing-in-go-example/example03_edges/ast"
	"github.com/awalterschulze/parsing-in-go-example/example03_edges/gen_parser/lexer"
	"github.com/awalterschulze/parsing-in-go-example/example03_edges/gen_parser/parser"
)

func TestPass(t *testing.T) {
	input := []byte("digraph Parsing { File -> Scanner\nScanner -> Parser }")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	got, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
	}
	want := &Graph{Name: "Parsing", Edges: []*Edge{{Src: "File", Dst: "Scanner"}, {Src: "Scanner", Dst: "Parser"}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected ast to be %v, but got %v", want, got)
	}
}

func TestFail(t *testing.T) {
	input := []byte("digraph { File }")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	_, err := p.Parse(lex)
	if err == nil {
		t.Fatal("expected error")
	}
}
