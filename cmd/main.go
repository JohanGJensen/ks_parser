package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	tokenFile, err := os.Open("tmp/tokens.txt")
	if err != nil {
		log.Println("could not open tokens file")
		return
	}
	defer tokenFile.Close()

	var tokens []string

	scanner := bufio.NewScanner(tokenFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		tokens = append(tokens, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file: ", err)
	}

	print(tokens)
}
