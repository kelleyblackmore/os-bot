// Package cli implements simple command handling for os-bot.
package cli

import (
	"bufio"
	"fmt"
	"os"
)

// Version of the os-bot application.
const Version = "0.1.0"

// RunInteractive starts an interactive shell allowing the user to type commands.
func RunInteractive() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Interactive mode. Type 'help' for commands. Type 'quit' to exit.")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if text == "quit" || text == "exit" {
			return
		}
		if text == "help" {
			fmt.Println("Commands: hello, version, help, quit")
			continue
		}
		HandleCommand(text)
	}
}

// HandleCommand executes a single command string.
func HandleCommand(cmd string) {
	switch cmd {
	case "hello":
		fmt.Println("Hello, world!")
	case "version":
		fmt.Printf("os-bot version %s\n", Version)
	default:
		fmt.Printf("unknown command: %s\n", cmd)
	}
}
