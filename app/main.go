package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

func main() {
	for {
		// Print the prompt
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo") {
			fmt.Printf("%s\n", command[5:])
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
