package tokens

func (t *RawTokens) GetParsedToken() (NodeType, string, error) {
	line := t.Current
	token, value, err := ParseRawTokenString(line)
	if err != nil {
		return "", "", err
	}

	return token, value, nil
}

func (t *RawTokens) IncrementToken() {
	t.Index += 1
	t.Current = t.Tokens[t.Index]
}
