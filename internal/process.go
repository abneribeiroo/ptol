package internal

import (
    "log"
    "os"
    "os/exec"
)

func startProcess(script string) *exec.Cmd {
    clearTerminal()

    var cmd *exec.Cmd
    switch {
    case script[len(script)-3:] == ".go":
        cmd = exec.Command("go", "run", script)
    case script[len(script)-3:] == ".js":
        cmd = exec.Command("node", script)
    default:
        log.Fatalf("Unsupported file type: %s", script)
    }

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        log.Fatalf("Error starting process: %v", err)
    }

    return cmd
}
