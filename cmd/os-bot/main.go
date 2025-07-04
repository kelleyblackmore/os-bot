package main

import (
	"flag"
	"fmt"

	"github.com/example/os-bot/internal/cli"
)

func main() {
	interactive := flag.Bool("interactive", false, "run in interactive mode")
	flag.Parse()

	if *interactive {
		cli.RunInteractive()
		return
	}

	if flag.NArg() == 0 {
		fmt.Println("Usage: os-bot [flags] <command>")
		fmt.Println("Commands: hello, version")
		return
	}

	cli.HandleCommand(flag.Arg(0))
}
