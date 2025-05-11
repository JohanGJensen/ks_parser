package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var refCount uint8

type RawTokens struct {
	Index   int
	Tokens  []string
	Current string
}

func (t *RawTokens) nextRawToken(index int) (string, error) {
	if len(t.Tokens) < index {
		return "", errors.New("index param exceeds length of raw tokens list")
	}

	t.Index = index
	t.Current = t.Tokens[index]

	return t.Current, nil
}

func (t *RawTokens) getCurrentToken() (NodeType, string) {
	line := t.Current
	token, value, err := parseRawTokenString(line)
	if err != nil {
		return "", ""
	}

	return token, value
}

func parseRawTokenString(t string) (NodeType, string, error) {
	splitToken := strings.Split(t, ",")
	token := NodeType(strings.Trim(strings.Replace(splitToken[0], "token:", "", -1), " "))
	value := strings.Trim(strings.Replace(splitToken[1], "value:", "", -1), " ")

	if token.isValidNodeType() {
		return token, value, nil
	}

	if token.isValidNodeTypeBracket() {
		return token, value, nil
	}

	return "", "", errors.New("could not split raw token string")
}

type SyntaxTreeArgument struct {
	Ref  uint8
	Name string
}

type SyntaxTreeConditional struct {
	Ref  uint8
	Name string
}

type SyntaxTreeObject struct {
	Typ          NodeType
	Name         string
	Ref          uint8
	Arguments    []SyntaxTreeArgument
	Conditionals []SyntaxTreeConditional
	Scope        []SyntaxTreeObject
}

type SyntaxTreeRoot struct {
	Typ   string
	Scope []SyntaxTreeObject
	Entry SyntaxTreeObject
}

type NodeType string

const (
	FUNCTION  NodeType = "Function"
	VARIABLE  NodeType = "VarName" // TODO -> should be variable
	OPERATOR  NodeType = "Operator"
	STATEMENT NodeType = "Statement"
	// brackets
	LEFTPAREN  NodeType = "ParentLeft"
	RIGHTPAREN NodeType = "ParentRight"
	LEFTBRACK  NodeType = "FunctionScopeStart"
	RIGHTBRACK NodeType = "FunctionScopeEnd"
)

func (n NodeType) isValidNodeType() bool {
	return n == FUNCTION || n == VARIABLE || n == OPERATOR || n == STATEMENT
}

func (n NodeType) isValidNodeTypeBracket() bool {
	return n == LEFTPAREN || n == RIGHTPAREN || n == LEFTBRACK || n == RIGHTBRACK
}

func (n NodeType) isNodeTypeFunction() bool {
	return n == FUNCTION
}

func (n NodeType) isNodeTypeVariable() bool {
	return n == VARIABLE
}

func main() {
	tokenFile, err := os.Open("tmp/tokens.txt")
	if err != nil {
		log.Println("could not open tokens file")
		return
	}
	defer tokenFile.Close()

	var tokenStruct RawTokens = RawTokens{Index: 0}

	scanner := bufio.NewScanner(tokenFile)
	for scanner.Scan() {
		line := scanner.Text()

		tokenStruct.Tokens = append(tokenStruct.Tokens, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file: ", err)
	}

	buildAST(tokenStruct)
}

// recursive descent parser (LL, TTB)
// lookup main first
func buildAST(tokenStruct RawTokens) {
	newAST := SyntaxTreeRoot{
		Typ:   "program",
		Scope: []SyntaxTreeObject{},
		// Entry -> find in range below
	}

	for i := 0; len(tokenStruct.Tokens) > i; i++ {
		rawToken, err := tokenStruct.nextRawToken(i)
		if err != nil {
			log.Fatal(err)
		}

		token, _, _ := parseRawTokenString(rawToken)

		if token.isNodeTypeFunction() {
			newAST.buildFunctionDeclarationScope(tokenStruct)
		}
	}

	log.Print(newAST)
}

func (ast *SyntaxTreeRoot) buildFunctionDeclarationScope(tokenStruct RawTokens) (SyntaxTreeObject, error) {
	token, _ := tokenStruct.getCurrentToken()
	refCount += 1
	newScopeObject := SyntaxTreeObject{
		Typ:          token,
		Ref:          refCount,
		Arguments:    []SyntaxTreeArgument{},
		Conditionals: []SyntaxTreeConditional{},
		Scope:        []SyntaxTreeObject{},
	}

	// set function name
	rawVariableToken, err := tokenStruct.nextRawToken(tokenStruct.Index + 1)
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	parsedToken, parsedValue, err := parseRawTokenString(rawVariableToken)
	if err != nil {
		return SyntaxTreeObject{}, err
	}

	if !parsedToken.isNodeTypeVariable() {
		return SyntaxTreeObject{}, errors.New("function: incorrect token type for function name")
	}

	newScopeObject.Name = parsedValue

	// build arguments
	rawLeftParamToken, err := tokenStruct.nextRawToken(tokenStruct.Index + 1)
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	_, parsedLeftParameterValue, _ := parseRawTokenString(rawLeftParamToken)
	if !strings.Contains(parsedLeftParameterValue, "(") {
		return SyntaxTreeObject{}, errors.New("syntax incorrect: expected (")
	}

	for i := tokenStruct.Index; len(tokenStruct.Tokens) > i; i++ {
		rawToken, err := tokenStruct.nextRawToken(i)
		if err != nil {
			return SyntaxTreeObject{}, errors.New("function: could not get next raw token")
		}

		// TODO: should append arguments

		parsedToken, parsedValue, _ := parseRawTokenString(rawToken)
		if strings.Contains(parsedValue, ")") {
			log.Println("closing argument: ", parsedToken)
			break
		}
	}

	// build scope
	rawLeftBracketToken, err := tokenStruct.nextRawToken(tokenStruct.Index + 1)
	if err != nil {
		return SyntaxTreeObject{}, errors.New("function: failed to get next token")
	}

	_, parsedLeftBracketValue, _ := parseRawTokenString(rawLeftBracketToken)
	if !strings.Contains(parsedLeftBracketValue, "{") {
		return SyntaxTreeObject{}, errors.New("syntax incorrect: expected {")
	}
	for i := tokenStruct.Index; len(tokenStruct.Tokens) > i; i++ {
		rawToken, err := tokenStruct.nextRawToken(i)
		if err != nil {
			return SyntaxTreeObject{}, errors.New("function: could not get next raw token")
		}

		// TODO: Function scope

		parsedToken, parsedValue, _ := parseRawTokenString(rawToken)
		if strings.Contains(parsedValue, "}") {
			log.Println("closing scope: ", parsedToken)
			break
		}
	}

	ast.Scope = append(ast.Scope, newScopeObject)
	return newScopeObject, nil
}
