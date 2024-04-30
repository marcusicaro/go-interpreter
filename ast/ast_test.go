package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: []rune("let")},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: []rune("myVar")},
					Value: []rune("myVar"),
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: []rune("anotherVar")},
					Value: []rune("anotherVar"),
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
