package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syntax_analyzer/internal/tokens"
	"syntax_analyzer/pkg/ast"
)

func main() {
	tokenFile, err := os.Open("tmp/tokens.txt")
	if err != nil {
		log.Println("could not open tokens file")
		return
	}
	defer tokenFile.Close()

	var tokenStruct tokens.RawTokens = tokens.RawTokens{Index: 0}

	scanner := bufio.NewScanner(tokenFile)
	for scanner.Scan() {
		line := scanner.Text()

		tokenStruct.Tokens = append(tokenStruct.Tokens, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file: ", err)
	}

	ast.BuildAST(tokenStruct)
}
