package lexer

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	NEWLINE TokenType = "NEWLINE"
	STRING  TokenType = "STRING"
	EOF     TokenType = "EOF"
	DISPLAY TokenType = "DISPLAY"
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}
