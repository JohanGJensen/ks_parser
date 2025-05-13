package ast

import "syntax_analyzer/internal/tokens"

type SyntaxTreeArgument struct {
	Ref  uint8
	Name string
}

type SyntaxTreeConditional struct {
	Ref  uint8
	Name string
}

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

type SyntaxTreeRoot struct {
	Typ   string
	Scope []SyntaxTreeObject
	Entry SyntaxTreeObject
}
