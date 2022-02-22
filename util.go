package main

import (
	"fmt"
)

var (
	commit  = ""
	version = "next"
)

func printVersion() {
	fmt.Printf("Clevyr Scaffold v%s", version)
	if commit != "" {
		fmt.Printf(" (%s)", commit)
	}
	fmt.Println()
}
