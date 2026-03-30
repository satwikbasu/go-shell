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
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the shell prompt
		fmt.Printf("$ ")
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
			echoPrint := strings.Join(parts[1:], " ")
			fmt.Printf("%s\n", echoPrint)
		
		// type command parse and eval
		} else if strings.HasPrefix(command, "type") {

			if len(parts) > 1 {
				argAfterType := parts[1] // fetching second word after type

				// check if type is followed by built-in shell command
				if slices.Contains(builtIn, argAfterType) {
					fmt.Printf("%s is a shell builtin\n", argAfterType)
				} else {
					pathValue := os.Getenv("PATH") // stores $PATH as a string
					directories := strings.Split(pathValue, string(os.PathListSeparator)) // splits $PATH directories into a slice
					// fmt.Printf("%#v\n", directories)
					// fmt.Printf("%T\n", pathValue)

					var execFound bool // to check if executable was found with right permissions

					for _, dir := range directories {
						fullPath := dir + "/" + argAfterType
						fileInfo, err := os.Stat(fullPath)
						// fmt.Printf("%#v\n", fullPath)
						if err == nil {
							mode := fileInfo.Mode()
							// fmt.Printf("%#v\n", mode)
							// Check if the found file is not a directory and has any execute permissions set (owner, group, or others)
							// mode.IsDir() returns true if the file is a directory, so !mode.IsDir() ensures it's a file
							// mode.Perm()&0111 checks if any of the execute bits are set (octal 0111 = execute for owner, group, or others)
							if !mode.IsDir() && (mode.Perm()&0111 != 0) {
								// Found an executable file in the directory
								fmt.Printf("%s is %s\n", argAfterType, fullPath)
								execFound = true
								break
							}
						}
					}
					if !execFound {
						fmt.Printf("%s: not found\n", argAfterType)
					}
				}
			}
		
		// invalid command handling
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}
