package iexec

import (
	"os"
	"os/exec"
)

func Command(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}