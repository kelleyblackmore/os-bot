# os-bot

os-bot aims to be an MCP-like agent for your operating system. The goal is a command-line assistant that can perform actions and offer recommendations to help you manage tasks. It can run interactively or execute a single command.

## Building

```bash
go build ./cmd/os-bot
```

## Running

To use interactive mode:

```
./cmd/os-bot/os-bot -interactive
```

Type `help` inside the prompt to see available commands.

You can also run a single command directly:

```
./os-bot <command>
```

