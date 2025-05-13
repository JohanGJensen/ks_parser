package ast

import (
	"errors"
	"log"
	"strings"
	t "syntax_analyzer/internal/tokens"
)

var refCount uint8

func (ast *SyntaxTreeRoot) buildVariableScope(tokenStruct t.RawTokens) {
	token, _ := tokenStruct.GetCurrentToken()
	refCount += 1
	newVariableObject := SyntaxTreeObject{
		Typ: token,
		Ref: refCount,
	}

	// VARIABLE_NAME token handling
	rawVariableNameToken, err := tokenStruct.NextRawToken(tokenStruct.Index + 1)
	if err != nil {
		// return SyntaxTreeObject{}, errors.New("function: failed to get next token")
		return
	}
	parsedVariableNameToken, parsedVariableNameValue, err := t.ParseRawTokenString(rawVariableNameToken)
	if err != nil {
		return
	}
	if !parsedVariableNameToken.IsNodeTypeVariableName() {
		return
	}
	newVariableObject.Name = parsedVariableNameValue

	// ASSIGN token handling
	rawAssignToken, err := tokenStruct.NextRawToken(tokenStruct.Index + 1)
	if err != nil {
		// return SyntaxTreeObject{}, errors.New("function: failed to get next token")
		return
	}
	parsedAssignToken, _, err := t.ParseRawTokenString(rawAssignToken)
	if err != nil {
		return
	}
	if parsedAssignToken != t.ASSIGN {
		return
	}

	// VARIABLE VALUE token handling
	for {
		rawToken, err := tokenStruct.NextRawToken(tokenStruct.Index + 1)
		if err != nil {
			// return SyntaxTreeObject{}, errors.New("function: failed to get next token")
			return
		}
		parsedToken, parsedValue, _ := t.ParseRawTokenString(rawToken)

		log.Println("identifier: ", parsedToken, parsedValue)
		if parsedToken.IsNodeTypeVariableValue() {
			newVariableObject.Value = parsedValue
			break
		}

	}

	// log.Println("variable value: ", parsedVariableValueToken, parsedVariableValueValue)
	log.Println(newVariableObject)

	ast.Scope = append(ast.Scope, newVariableObject)
}

func (ast *SyntaxTreeRoot) buildFunctionDeclarationScope(tokenStruct t.RawTokens) (SyntaxTreeObject, error) {
	token, _ := tokenStruct.GetCurrentToken()
	refCount += 1
	newScopeObject := SyntaxTreeObject{
		Typ:          token,
		Ref:          refCount,
		Arguments:    []SyntaxTreeArgument{},
		Conditionals: []SyntaxTreeConditional{},
		Scope:        []SyntaxTreeObject{},
	}

	// set function name
	rawVariableToken, err := tokenStruct.NextRawToken(tokenStruct.Index + 1)
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	parsedToken, parsedValue, err := t.ParseRawTokenString(rawVariableToken)
	if err != nil {
		return SyntaxTreeObject{}, err
	}

	if !parsedToken.IsNodeTypeVariableName() {
		return SyntaxTreeObject{}, errors.New("function: incorrect token type for function name")
	}

	newScopeObject.Name = parsedValue

	// build arguments
	rawLeftParamToken, err := tokenStruct.NextRawToken(tokenStruct.Index + 1)
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	_, parsedLeftParameterValue, _ := t.ParseRawTokenString(rawLeftParamToken)
	if !strings.Contains(parsedLeftParameterValue, "(") {
		return SyntaxTreeObject{}, errors.New("syntax incorrect: expected (")
	}

	for i := tokenStruct.Index; len(tokenStruct.Tokens) > i; i++ {
		rawToken, err := tokenStruct.NextRawToken(i)
		if err != nil {
			return SyntaxTreeObject{}, errors.New("function: could not get next raw token")
		}

		// TODO: should append arguments

		parsedToken, parsedValue, _ := t.ParseRawTokenString(rawToken)
		if strings.Contains(parsedValue, ")") {
			log.Println("closing argument: ", parsedToken)
			break
		}
	}

	// build scope
	rawLeftBracketToken, err := tokenStruct.NextRawToken(tokenStruct.Index + 1)
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	_, parsedLeftBracketValue, _ := t.ParseRawTokenString(rawLeftBracketToken)
	if !strings.Contains(parsedLeftBracketValue, "{") {
		return SyntaxTreeObject{}, errors.New("syntax incorrect: expected {")
	}
	for i := tokenStruct.Index; len(tokenStruct.Tokens) > i; i++ {
		rawToken, err := tokenStruct.NextRawToken(i)
		if err != nil {
			return SyntaxTreeObject{}, errors.New("function: could not get next raw token")
		}

		// TODO: Function scope

		parsedToken, parsedValue, _ := t.ParseRawTokenString(rawToken)
		if strings.Contains(parsedValue, "}") {
			log.Println("closing scope: ", parsedToken)
			break
		}
	}

	ast.Scope = append(ast.Scope, newScopeObject)
	return newScopeObject, nil
}
