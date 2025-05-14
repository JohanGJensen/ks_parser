package ast

import (
	"errors"
	t "syntax_analyzer/internal/tokens"
)

/*
Builder function that creates a syntax tree node variable
*/
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
		tokenStruct.IncrementToken()
		parsedToken, parsedValue, _ := tokenStruct.GetParsedToken()

		if parsedToken.IsNodeTypeVariableValue() {
			if parsedToken == t.VARIABLE_NAME {
				execution, err := ast.buildFunctionExecution(*tokenStruct)
				if err != nil {
					return err
				}

				newVariableObject.Value = execution
				break
			}

			newVariableObject.Value = parsedValue
			break
		}
	}

	ast.Scope = append(ast.Scope, newVariableObject)
	return nil
}

/*
Helper function to variable scope builder function.
checks the current parsed token if it the correct anticipated terminal grammar "assign".
*/
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
