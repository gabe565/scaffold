package main

import (
	"github.com/clevyr/scaffold/cmd"
	"os"
)

//go:generate go run internal/generators/template_embed.go

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
