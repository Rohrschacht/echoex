package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	timestamp := false

	if len(os.Args) > 1 && os.Args[1] == "-t" {
		args = os.Args[2:]
		timestamp = true
	}

	fmt.Printf("`%s`", strings.Join(args, " "))
	if timestamp {
		fmt.Printf(" at \"%s\"", time.Now().Local().Format("Mon Jan 2 2006 15:04:05"))
	}
	fmt.Printf(":\n")

	if len(args) == 0 {
		fmt.Printf("no program specified!\n")
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			fmt.Printf("echoex encountered an unexpected error: %v\n", err)
			os.Exit(-1)
		}
	}
}
