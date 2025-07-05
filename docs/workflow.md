# os-bot Workflow

The diagram below illustrates how `os-bot` operates when using interactive mode with models served by Ollama.

```mermaid
sequenceDiagram
    participant U as User
    participant C as CLI
    participant M as Ollama Model
    participant OS as Operating System

    U->>C: Prompt in interactive CLI
    C->>M: Send prompt to model
    M->>C: Suggest actions or checks
    C->>OS: Perform system changes or checks
    OS->>C: Return command results
    C->>U: Display output to user
```

In this workflow, the user interacts with the command-line interface, which communicates with a model served by Ollama. The model suggests tasks to perform on the operating system, and the CLI executes them, displaying any results back to the user.
