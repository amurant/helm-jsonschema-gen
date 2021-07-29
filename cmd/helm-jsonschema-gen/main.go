package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := NewCmdJsonSchemaGen()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
