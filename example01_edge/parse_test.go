package example01_edge

import (
	"testing"

	"github.com/awalterschulze/parsing-in-go-example/example01_edge/gen_parser/lexer"
	"github.com/awalterschulze/parsing-in-go-example/example01_edge/gen_parser/parser"
)

func TestPass(t *testing.T) {
	input := []byte("File -> Scanner")
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	_, err := p.Parse(lex)
	if err != nil {
		t.Fatal(err)
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
