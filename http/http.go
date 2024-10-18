package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func request() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lw := logWriter{}

	_, _ = io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
