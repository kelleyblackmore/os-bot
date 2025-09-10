// Package cli implements simple command handling for os-bot.
package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/example/os-bot/internal/ollama"
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
			fmt.Println("Commands: hello, version, ask, help, quit")
			continue
		}
		HandleCommand(text)
	}
}

// HandleCommand executes a single command string.
func HandleCommand(cmd string) {
	parts := strings.SplitN(cmd, " ", 2)
	command := parts[0]
	var arg string
	if len(parts) > 1 {
		arg = parts[1]
	}

	switch command {
	case "hello":
		fmt.Println("Hello, world!")
	case "version":
		fmt.Printf("os-bot version %s\n", Version)
	case "ask":
		if arg == "" {
			fmt.Println("Usage: ask <prompt>")
			return
		}
		response, err := ollama.Ask(arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(response)
	default:
		fmt.Printf("unknown command: %s\n", cmd)
	}
}
