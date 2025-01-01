package internal

import (
    "os"
    "os/exec"
)

func clearTerminal() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
