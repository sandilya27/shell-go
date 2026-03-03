package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		parts := strings.Fields(command)

		command = parts[0]

		var args string
		if len(parts) > 1 {
			args = strings.Join(parts[1:], " ")
		}

		if command == "exit" {
			os.Exit(0)
		}

		if command == "echo" {
			fmt.Printf("%s\n", args)
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}

}
