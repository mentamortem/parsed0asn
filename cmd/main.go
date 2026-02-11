package main

import (
	"fmt"
	"os"

	parser "github.com/mentamortem/parsed0asn/internal/parser"
)

func main() {
	path := os.Args[1]
	raw, err := parser.ProcessFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(raw[0:10000])
}
