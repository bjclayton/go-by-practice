package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
