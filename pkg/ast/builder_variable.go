package ast

import (
	"errors"
	t "syntax_analyzer/internal/tokens"
)

func (ast *SyntaxTreeRoot) buildVariableScope(tokenStruct *t.RawTokens) error {
	token, _, _ := tokenStruct.GetParsedToken()
	refCount += 1
	newVariableObject := SyntaxTreeObject{
		Typ: token,
		Ref: refCount,
	}

	// check that the next token terminal grammar is a valid variable name
	tokenStruct.IncrementToken()
	errVarName := newVariableObject.setVarName(tokenStruct)
	if errVarName != nil {
		return errVarName
	}

	// check that the next token terminal grammar is a valid assign operator
	tokenStruct.IncrementToken()
	errAssign := checkTokenIsAssign(tokenStruct)
	if errAssign != nil {
		return errAssign
	}

	// check that the next token terminal grammar is a valid variable value
	for {
		parsedToken, parsedValue, _ := tokenStruct.GetNextParsedToken()

		if parsedToken.IsNodeTypeVariableValue() {
			newVariableObject.Value = parsedValue
			break
		}
	}

	ast.Scope = append(ast.Scope, newVariableObject)
	return nil
}

func checkTokenIsAssign(tokenStruct *t.RawTokens) error {
	token, _, err := tokenStruct.GetParsedToken()
	if err != nil {
		return err
	}

	if token != t.ASSIGN {
		return errors.New("assign token does not match terminal grammar")
	}

	return nil
}
