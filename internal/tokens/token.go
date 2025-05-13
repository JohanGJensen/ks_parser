package tokens

import "errors"

func (t *RawTokens) nextRawToken() (string, error) {
	incrementedIndex := t.Index + 1

	if len(t.Tokens) < incrementedIndex {
		return "", errors.New("index param exceeds length of raw tokens list")
	}

	t.Index = incrementedIndex
	t.Current = t.Tokens[incrementedIndex]

	return t.Current, nil
}

func (t *RawTokens) GetParsedToken() (NodeType, string, error) {
	line := t.Current
	token, value, err := ParseRawTokenString(line)
	if err != nil {
		return "", "", err
	}

	return token, value, nil
}

func (t *RawTokens) GetNextParsedToken() (NodeType, string, error) {
	rawToken, err := t.nextRawToken()
	if err != nil {
		return "", "", err
	}
	token, value, err := ParseRawTokenString(rawToken)
	if err != nil {
		return "", "", err
	}

	return token, value, nil
}

func (t *RawTokens) IncrementToken() {
	t.Index += 1
	t.Current = t.Tokens[t.Index]
}
