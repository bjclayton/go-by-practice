package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//io.Copy(os.Stdout, res.Body)

	bs := make([]byte, 1024)
	for {
		n, err := res.Body.Read(bs)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}

		fmt.Print(string(bs[:n]))
	}
}
