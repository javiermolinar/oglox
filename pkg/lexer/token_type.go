package lexer

type TokenType int

const (
	LEFT_PARENT TokenType = iota
	RIGHT_PARENT
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two characters tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

var TokenStrings = []string{
	"LEFT_PARENT", "RIGHT_PARENT", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR",
	"BANG", "BANG_EQUAL", "EQUAL", "EQUAL_EQUAL", "GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL", "IDENTIFIER", "STRING", "NUMBER",
	"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR", "PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR",
	"WHILE", "EOF",
}

func (t TokenType) String() string {
	return TokenStrings[t]
}
