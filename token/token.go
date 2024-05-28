package token

type TokenType []rune

type Token struct {
	Type    TokenType
	Literal []rune
}

var (
	ILLEGAL = []rune("ILLEGAL")
	EOF     = []rune("EOF")

	// Identifiers + literals
	IDENT = []rune("IDENT") // add, foobar, x, y, ...
	INT   = []rune("INT")   // 123456

	// Operators
	ASSIGN   = []rune("=")
	PLUS     = []rune("+")
	MINUS    = []rune("-")
	BANG     = []rune("!")
	ASTERISK = []rune("*")
	SLASH    = []rune("/")

	LT = []rune("<")
	GT = []rune(">")
	// Delimiters
	COMMA     = []rune(",")
	SEMICOLON = []rune(";")
	LPAREN    = []rune("(")
	RPAREN    = []rune(")")
	LBRACE    = []rune("{")
	RBRACE    = []rune("}")
	// Keywords
	FUNCTION = []rune("FUNCTION")
	LET      = []rune("LET")
	TRUE     = []rune("TRUE")
	FALSE    = []rune("FALSE")
	IF       = []rune("IF")
	ELSE     = []rune("ELSE")
	RETURN   = []rune("RETURN")

	EQ     = []rune("==")
	NOT_EQ = []rune("!=")

	GREATERTHANOREQUALTO = []rune(">=")
	LESSTHANOREQUALTO    = []rune("<=")

	STRING = []rune("STRING")
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident []rune) TokenType {
	if tok, ok := keywords[string(ident)]; ok {
		return tok
	}
	return IDENT
}
