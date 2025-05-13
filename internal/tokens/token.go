package tokens

import "errors"

func (t *RawTokens) NextRawToken(index int) (string, error) {
	if len(t.Tokens) < index {
		return "", errors.New("index param exceeds length of raw tokens list")
	}

	t.Index = index
	t.Current = t.Tokens[index]

	return t.Current, nil
}

func (t *RawTokens) GetCurrentToken() (NodeType, string) {
	line := t.Current
	token, value, err := ParseRawTokenString(line)
	if err != nil {
		return "", ""
	}

	return token, value
}
