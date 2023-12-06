package main

import (
	"log"

	"github.com/schema-creator/services/scml/lexer"
	"github.com/schema-creator/services/scml/token"
)

func main() {
	file := `
		Table users {
			user_id int [pk, increment]
			user_name varchar [nn]
			user_email varchar [nn, uq]
			user_password varchar [nn]
		}	
	`
	lex := lexer.New(file)
	for {
		tok := lex.NextToken()
		log.Println(tok)
		if tok.Type == token.EOF {
			break
		}
	}
}
