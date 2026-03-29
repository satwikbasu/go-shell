package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

func main() {
	// Print the prompt
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	command = strings.TrimSpace(command)
	fmt.Printf("%s: command not found\n", command)
}
