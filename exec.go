package main

import (
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	if args[0] == "cd" {
		return os.Chdir(args[1])
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
