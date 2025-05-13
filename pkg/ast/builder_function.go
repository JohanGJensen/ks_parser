package ast

import (
	"errors"
	"log"
	"strings"
	t "syntax_analyzer/internal/tokens"
)

func (ast *SyntaxTreeRoot) buildFunctionDeclarationScope(tokenStruct t.RawTokens) (SyntaxTreeObject, error) {
	token, _, _ := tokenStruct.GetParsedToken()
	refCount += 1
	newScopeObject := SyntaxTreeObject{
		Typ:          token,
		Ref:          refCount,
		Arguments:    []SyntaxTreeArgument{},
		Conditionals: []SyntaxTreeConditional{},
		Scope:        []SyntaxTreeObject{},
	}

	// set function name
	parsedToken, parsedValue, _ := tokenStruct.GetNextParsedToken()
	if !parsedToken.IsNodeTypeVariableName() {
		return SyntaxTreeObject{}, errors.New("function: incorrect token type for function name")
	}
	newScopeObject.Name = parsedValue

	// build arguments
	_, parsedLeftParamValue, _ := tokenStruct.GetNextParsedToken()
	if !strings.Contains(parsedLeftParamValue, "(") {
		return SyntaxTreeObject{}, errors.New("syntax incorrect: expected (")
	}

	for i := tokenStruct.Index; len(tokenStruct.Tokens) > i; i++ {
		parsedToken, parsedValue, err := tokenStruct.GetParsedToken()
		if err != nil {
			return SyntaxTreeObject{}, errors.New("function: could not get next raw token")
		}

		// TODO: should append arguments

		if strings.Contains(parsedValue, ")") {
			log.Println("closing argument: ", parsedToken)
			break
		}
	}

	// build scope
	_, parsedValue, err := tokenStruct.GetNextParsedToken()
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	if !strings.Contains(parsedValue, "{") {
		return SyntaxTreeObject{}, errors.New("syntax incorrect: expected {")
	}
	for i := tokenStruct.Index; len(tokenStruct.Tokens) > i; i++ {
		parsedToken, parsedValue, err := tokenStruct.GetParsedToken()
		if err != nil {
			return SyntaxTreeObject{}, errors.New("function: could not get next raw token")
		}

		// TODO: Function scope

		if strings.Contains(parsedValue, "}") {
			log.Println("closing scope: ", parsedToken)
			break
		}
	}

	ast.Scope = append(ast.Scope, newScopeObject)
	return newScopeObject, nil
}
