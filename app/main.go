package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"slices"
)

func main() {
	for {
		// Print the shell prompt
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command) // remove trailing whitespaces
		builtIn := []string{"echo", "exit", "type"} // slice to check for built-in shell commands

		// parts stores all words in the command separately 
		parts := strings.Fields(command)

		// exit command parse and eval
		if command == "exit" {
			break
		
		// echo command parse and eval
		} else if strings.HasPrefix(command, "echo") {
			fmt.Printf("%s\n", command[5:])
		
		// type command parse and eval
		} else if strings.HasPrefix(command, "type") {

			if len(parts) > 1 {
				argAfterType := parts[1] // fetching second word after type

				// check if type is followed by built-in shell command
				if slices.Contains(builtIn, argAfterType) {
					fmt.Printf("%s is a shell builtin\n", argAfterType)
				} else {
					fmt.Printf("%s: not found\n", argAfterType)
				}
			}
		
		// invalid command handling
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}
