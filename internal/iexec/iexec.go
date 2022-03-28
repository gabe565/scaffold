package iexec

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Builder struct {
	cmd    []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewBuilder(p ...string) *Builder {
	return &Builder{
		cmd:    p,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (b *Builder) Append(p ...string) *Builder {
	b.cmd = append(b.cmd, p...)
	return b
}

func (b Builder) Run() error {
	cmd := exec.Command(b.cmd[0], b.cmd[1:]...)
	fmt.Printf("+ %s\n", cmd)
	// Disable Telescope by default
	cmd.Env = append(os.Environ(), "TELESCOPE_ENABLED=false")
	cmd.Stdin = b.Stdin
	cmd.Stdout = b.Stdout
	cmd.Stderr = b.Stderr
	return cmd.Run()
}
