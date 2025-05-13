package ast

import (
	"log"
	"syntax_analyzer/internal/tokens"
)

// recursive descent parser (LL, TTB)
// lookup main first
func BuildAST(tokenStruct tokens.RawTokens) {
	newAST := SyntaxTreeRoot{
		Typ:   "program",
		Scope: []SyntaxTreeObject{},
		// Entry -> find in range below
	}

	for i := 0; len(tokenStruct.Tokens) > i; i++ {
		rawToken, err := tokenStruct.NextRawToken(i)
		if err != nil {
			log.Fatal(err)
		}

		token, _, _ := tokens.ParseRawTokenString(rawToken)

		if token.IsNodeTypeFunction() {
			newAST.buildFunctionDeclarationScope(tokenStruct)
		}

		if token.IsNodeTypeVariable() {
			newAST.buildVariableScope(tokenStruct)
		}
	}

	log.Print(newAST)
}
