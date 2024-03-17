package token

const (
	ILLEGAL = "ILLEGAL" // Token or character that we don't know about
	EOF     = "EOF"     // End Of File so that the parser can stop parsing

	// Identifiers (or variable names) + literals (or values)
	IDENT = "IDENT" // x, y, etc
	INT   = "INT"   // 1, 2, 3, etc

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters and special characters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType // Stores the type of token. Eg: number, special characters, etc
	Literal string    // Stores the actual value of the token. Eg: number value of 5 or 10, { or }
}
