package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"slices"
)

var builtIn []string = []string{"echo", "exit", "type"} // slice to check for built-in shell commands

// converts the raw terminal input into a slice
func command_to_list (inputString string) []string {
	command := strings.TrimSpace(inputString) // remove trailing whitespaces
	parts := strings.Fields(command) // stores all words in the command separately
	return parts
}

// prints output for invalid command
func print_invalid(command string) {
	fmt.Printf("%s: not found\n", command)
}

// checks if the command is builtin and has enough arguments
func check_for_builtin(commandList []string) {
	isBuiltin := slices.Contains(builtIn, commandList[0])
	if isBuiltin && len(commandList) > 1 {
		run_builtin(commandList)
	} else {
		// checks if there are no arguments at all
		if isBuiltin && len(commandList) == 1 {
			fmt.Printf("%s: too few arguments\n", commandList[0])
		} else {
			fmt.Printf("%s: not found\n", commandList[0])
		}
	}
}

// runs built in command
// TODO: refactor with switch case or better alternatives
func run_builtin(commandList []string) {

	// echo command parse and eval
	if commandList[0] == "echo" {
		echoPrint := strings.Join(commandList[1:], " ")
		fmt.Printf("%s\n", echoPrint)
	} else if commandList[0] == "type" {

		// type command parse and eval
		if len(commandList) > 1 {
			argAfterType := commandList[1] // fetching second word after type

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
					// fmt.Printf("%s: not found\n", argAfterType)
					print_invalid(argAfterType)
				}
			}
		}
	
	// invalid command handling
	} else {
		print_invalid(commandList[0])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the shell prompt
		fmt.Printf("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		commandList := command_to_list(input)
		
		// exit command parse and eval
		if commandList[0] == "exit" {
			break
		} else {
			check_for_builtin(commandList)
		}
	}
}
