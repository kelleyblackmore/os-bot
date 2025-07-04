package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
)

const version = "0.1.0"

func main() {
    interactive := flag.Bool("interactive", false, "run in interactive mode")
    flag.Parse()

    if *interactive {
        runInteractive()
        return
    }

    if flag.NArg() == 0 {
        fmt.Println("Usage: os-bot [flags] <command>")
        fmt.Println("Commands: hello, version")
        return
    }

    handleCommand(flag.Arg(0))
}

func runInteractive() {
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
        handleCommand(text)
    }
}

func handleCommand(cmd string) {
    switch cmd {
    case "hello":
        fmt.Println("Hello, world!")
    case "version":
        fmt.Printf("os-bot version %s\n", version)
    default:
        fmt.Printf("unknown command: %s\n", cmd)
    }
}

