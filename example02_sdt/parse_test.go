package example02_sdt

import (
	"reflect"
	"testing"

	. "github.com/awalterschulze/parsing-in-go-example/example02_sdt/ast"
	"github.com/awalterschulze/parsing-in-go-example/example02_sdt/gen_parser/lexer"
	"github.com/awalterschulze/parsing-in-go-example/example02_sdt/gen_parser/parser"
)

func TestPass(t *testing.T) {
	input := []byte("File -> Scanner")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	got, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
	}
	want := &Edge{Src: "File", Dst: "Scanner"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected ast to be %v, but got %v", want, got)
	}
}

func TestFail(t *testing.T) {
	input := []byte("File -> Scanner\nScanner -> Parser")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	_, err := p.Parse(lex)
	if err == nil {
		t.Fatal("expected error")
	}
}
