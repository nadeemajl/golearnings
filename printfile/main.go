package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote ", len(bs), " bytes")
	return len(bs), nil
}
