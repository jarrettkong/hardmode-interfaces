package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
