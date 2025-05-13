package tokens

type NodeType string

type RawTokens struct {
	Index   int
	Tokens  []string
	Current string
}
