package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func commandParsing(command string) (string, []string) {
	command = strings.TrimSpace(command)

	parts := strings.Fields(command)

	if len(parts) == 0 {
		return "", nil
	}

	return parts[0], parts[1:]
}

func handleTypePrint(command string) {
	switch command {
	case "echo", "exit", "type", "pwd":
		fmt.Printf("%s is a shell builtin\n", command)
	default:
		path, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("%s: not found\n", command)
			return
		}
		fmt.Printf("%s is %s\n", command, path)
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
			handleTypePrint(args[0])
		case "echo":
			fmt.Printf("%s\n", strings.Join(args, " "))
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(dir)
		case "exit":
			os.Exit(0)
		default:
			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			err := cmd.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}

}
