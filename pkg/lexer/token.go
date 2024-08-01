package lexer

import "fmt"

type token struct {
	tokenType TokenType
	lexeme    string
	literal   string
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal string, line int) *token {
	return &token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t token) String() string {
	return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, t.literal)
}
