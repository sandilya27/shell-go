package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandParsing(command string) (string, string) {
	command = strings.TrimSpace(command)

	parts := strings.Fields(command)

	command = parts[0]

	var args string
	if len(parts) > 1 {
		args = strings.Join(parts[1:], " ")
	}

	return command, args
}

func handleTypePrint(command string) {
	switch command {
	case "echo":
		fmt.Printf("%s: is a shell builtin\n", command)
	case "exit":
		fmt.Printf("%s: is a shell builtin\n", command)
	case "type":
		fmt.Printf("%s: is a shell builtin\n", command)
	default:
		fmt.Printf("%s: not found\n", command)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		command, args := commandParsing(command)

		switch command {
		case "type":
			cmd, _ := commandParsing(args)
			handleTypePrint(cmd)
		case "echo":
			fmt.Printf("%s\n", args)
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}

}
