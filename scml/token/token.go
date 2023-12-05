package token

import "strings"

// TokenType は「整数」や「閉じ括弧」などを区別するもの
type TokenType string

// Token はTokenTypeとリテラルを保持する型
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL はトークンや文字が未知であることを示す
	ILLEGAL = "ILLEGAL"
	// EOF はファイル終端。構文解析器にここで停止してよいと伝える
	EOF = "EOF"

	// IDENT は識別子（add, foobar, x, y, ...)
	IDENT = "IDENT"

	TYPES = "TYPES"

	COMMENT = "//"
	// LBRACE は左中括弧
	LBRACE = "{"
	// RBRACE は右中括弧
	RBRACE = "}"
	// LBRACKET は左大括弧
	LBRACKET = "["
	// RBRACKET は右大括弧
	RBRACKET = "]"
	// COLON はコロン（:）
	COLON = ":"

	PRIMARY_KEY = "pk"
	UNIQUE      = "uq"
	NOT_NULL    = "nn"
	INCREMENT   = "incliment"
	DEFAULT     = "df"

	TABLE = "TABLE"
	ENUM  = "ENUM"
	INDEX = "INDEX"
)

var keywords = map[string]TokenType{
	"table": TABLE,
	"enum":  ENUM,
	"index": INDEX,

	"pk":        PRIMARY_KEY,
	"uq":        UNIQUE,
	"nn":        NOT_NULL,
	"incliment": INCREMENT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}
