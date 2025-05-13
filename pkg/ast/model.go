package ast

type SyntaxTreeArgument struct {
	Ref  uint8
	Name string
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
