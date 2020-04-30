package main

import (
	"os"
	"os/exec"
	"strings"
)

var exit = os.Exit

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir(os.Getenv("HOME"))
		}
		return os.Chdir(args[1])
	case "exit":
		exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
