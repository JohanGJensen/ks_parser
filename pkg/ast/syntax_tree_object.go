package ast

import (
	"errors"
	"syntax_analyzer/internal/tokens"
)

type SyntaxTreeObject struct {
	Typ          tokens.NodeType
	Name         string
	Ref          uint8
	Arguments    []SyntaxTreeArgument
	Conditionals []SyntaxTreeConditional
	Scope        []SyntaxTreeObject
	Value        interface{} // basically "any" type
	// VarType string | int
}

func (sto *SyntaxTreeObject) setVarName(tokenStruct *tokens.RawTokens) error {
	token, value, err := tokenStruct.GetParsedToken()
	if err != nil {
		return err
	}

	if !token.IsNodeTypeVariableName() {
		return errors.New("variable name token does not match terminal grammar")
	}

	sto.Name = value
	return nil
}
