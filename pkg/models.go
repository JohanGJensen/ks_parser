package internal

type SyntaxTreeArgument struct {
	Ref  uint8
	Name string
}

type SyntaxTreeConditional struct {
	Ref  uint8
	Name string
}

type VariableUnion interface {
	string | int
}

type SyntaxTreeObject struct {
	Typ          NodeType
	Name         string
	Ref          uint8
	Arguments    []SyntaxTreeArgument
	Conditionals []SyntaxTreeConditional
	Scope        []SyntaxTreeObject
	// VarType string | int
}

type SyntaxTreeRoot struct {
	Typ   string
	Scope []SyntaxTreeObject
	Entry SyntaxTreeObject
}

type NodeType string
