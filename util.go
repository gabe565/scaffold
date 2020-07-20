package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	commit = ""
	version = "next"
)

func printVersion() {
	fmt.Printf("Clevyr Scaffold v%s", version)
	if commit != "" {
		fmt.Printf(" (%s)", commit)
	}
	fmt.Println()
}

func interactiveCommand(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}