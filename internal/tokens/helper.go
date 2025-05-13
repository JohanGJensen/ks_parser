package tokens

import (
	"errors"
	"strings"
)

func ParseRawTokenString(t string) (NodeType, string, error) {
	splitToken := strings.SplitN(t, ",", 2)
	token := NodeType(strings.Trim(strings.Replace(splitToken[0], "token:", "", -1), " "))
	value := strings.Trim(strings.Replace(splitToken[1], "value:", "", -1), " ")

	if token.IsValidNodeType() {
		return token, value, nil
	}

	if token.IsValidNodeTypeBracket() {
		return token, value, nil
	}

	if token.IsNodeTypeVariableValue() {
		return token, value, nil
	}

	if token.IsEndOfFile() {
		return token, value, nil
	}

	return "", "", errors.New("could not split raw token string")
}
