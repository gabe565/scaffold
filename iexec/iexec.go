package iexec

import (
	"fmt"
	"os"
	"os/exec"
)

func Command(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	fmt.Printf("+ %s\n", cmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}