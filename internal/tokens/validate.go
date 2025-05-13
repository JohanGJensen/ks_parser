package tokens

const (
	FUNCTION          NodeType = "Function"
	VARIABLE_NAME     NodeType = "VarName" // TODO -> should be variable
	VARIABLE_STRING   NodeType = "StringVar"
	VARIABLE_INTEGER  NodeType = "IntVar"
	STRING_IDENTIFIER NodeType = "StringIdentifier"
	INT_IDENTIFIER    NodeType = "NumberIdentifier"
	// STATEMENT        NodeType = "Statement"
	// brackets
	LEFTPAREN  NodeType = "ParentLeft"
	RIGHTPAREN NodeType = "ParentRight"
	LEFTBRACK  NodeType = "FunctionScopeStart"
	RIGHTBRACK NodeType = "FunctionScopeEnd"
	// Operators
	ASSIGN NodeType = "Assign"
	EOF    NodeType = "EOF"
)

func (n NodeType) IsValidNodeType() bool {
	return (n == FUNCTION ||
		n == VARIABLE_NAME ||
		n == VARIABLE_STRING ||
		n == VARIABLE_INTEGER ||
		n == ASSIGN)
}

func (n NodeType) IsValidNodeTypeBracket() bool {
	return n == LEFTPAREN || n == RIGHTPAREN || n == LEFTBRACK || n == RIGHTBRACK
}

func (n NodeType) IsNodeTypeFunction() bool {
	return n == FUNCTION
}

func (n NodeType) IsNodeTypeVariableName() bool {
	return n == VARIABLE_NAME
}

func (n NodeType) IsNodeTypeVariableValue() bool {
	return n == STRING_IDENTIFIER || n == INT_IDENTIFIER || n == VARIABLE_NAME
}

func (n NodeType) IsNodeTypeVariable() bool {
	return n == VARIABLE_STRING || n == VARIABLE_INTEGER
}

func (n NodeType) IsEndOfFile() bool {
	return n == EOF
}
