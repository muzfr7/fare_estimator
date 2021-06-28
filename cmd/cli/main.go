package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "", "-file paths.csv")
	flag.Parse()
	if *file == "" {
		fmt.Println("file parameter is required.")
		os.Exit(1)
	}
}
