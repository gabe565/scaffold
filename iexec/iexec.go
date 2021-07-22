package iexec

import (
	"fmt"
	"os"
	"os/exec"
)

func Command(name string, args ...string) (err error) {
	cmd := exec.Command(name, args...)
	fmt.Printf("+ %s\n", cmd)
	// Disable Telescope by default
	cmd.Env = append(os.Environ(), "TELESCOPE_ENABLED=false")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}
