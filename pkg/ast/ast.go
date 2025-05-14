package ast

import (
	"log"
	"syntax_analyzer/internal/tokens"
)

var refCount uint8

// recursive descent parser (LL, TTB)
// lookup main first
func BuildAbstractSyntaxTree(tokenStruct *tokens.RawTokens) {
	newAST := SyntaxTreeRoot{
		Typ:   "program",
		Scope: []SyntaxTreeObject{},
		// Entry -> find in range below
	}

	tokenStruct.Current = tokenStruct.Tokens[0]

	for {
		token, _, _ := tokens.ParseRawTokenString(tokenStruct.Current)
		if token.IsEndOfFile() {
			log.Default().Print("end of file reached")
			break
		}

		// TODO: fix/complete function scope builder
		// if token.IsNodeTypeFunction() {
		// 	newAST.buildFunctionDeclarationScope(tokenStruct)
		// }

		if token.IsNodeTypeVariable() {
			newAST.buildVariableScope(tokenStruct)
		}

		tokenStruct.IncrementToken()
	}

	log.Println(newAST)
}
