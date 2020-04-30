package main

import (
	"os"
	"testing"
)

func TestExecInput(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
		err  error
	}{
		{
			name: "cmd: ls",
			in:   "ls",
		},
		{
			name: "cmd: ls -l",
			in:   "ls -l",
		},
		{
			name: "cmd: cd",
			in:   "cd",
		},
		{
			name: "cmd: cd /",
			in:   "cd /",
			out:  "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := execInput(tt.in); err != tt.err {
				t.Errorf("execInput(%q) = %v; want %v", tt.in, err, tt.err)
			}

			if tt.out != "" {
				out, err := os.Getwd()
				if err != nil {
					t.Errorf("execInput(%q) = %v; want %v", tt.in, err, tt.err)
				}
				if out != tt.out {
					t.Errorf("execInput(%q) = %v; want %v", tt.in, out, tt.out)
				}
			}
		})
	}
}

func TestExecInputExit(t *testing.T) {

	exited := exit
	defer func() { exit = exited }()

	var got int
	newExit := func(code int) {
		got = code
	}

	exit = newExit
	execInput("exit")
	if want := 0; got != want {
		t.Errorf("got exit code %d; want %d", got, want)
	}
}
