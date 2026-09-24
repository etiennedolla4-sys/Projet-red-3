package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	var command *exec.Cmd

	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/c", "cls")
	} else {
		command = exec.Command("clear")
	}

	command.Stdout = os.Stdout
	command.Run()
}
