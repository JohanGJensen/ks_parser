package ast

type SyntaxTreeArgument struct {
	Ref   uint8
	Value string
}

// function = terminal symbol, execution = non terminal symbol
type SyntaxTreeFunctionExecution struct {
	Typ       string
	Name      string
	Arguments []SyntaxTreeArgument
}

type SyntaxTreeConditional struct {
	Ref  uint8
	Name string
}

type SyntaxTreeRoot struct {
	Typ   string
	Scope []SyntaxTreeObject
	Entry SyntaxTreeObject
}
