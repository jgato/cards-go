package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		f, err := os.Open(args[1])
		if err != nil {
			fmt.Println("error reading the file")
			os.Exit(1)
		}
		nBytes, err := io.Copy(os.Stdin, f)
		if err != nil {
			fmt.Println("Error printing file")
			os.Exit(1)
		}
		fmt.Printf("read %d bytes\n", nBytes)
	}
}
