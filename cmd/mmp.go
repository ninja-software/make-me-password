package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	mmp "github.com/ninja-software/make-me-password"
)

func main() {
	format := flag.String("format", "base64", "output format. choices: base64, hex")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("generate password for db use")
		fmt.Printf("usage: %s abc\n", path.Base(os.Args[0]))
		fmt.Printf("usage: %s -format=base64 abc\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	in := args[0]
	out, err := mmp.HashPassword(in, *format)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(2)
	}

	fmt.Println(out)
}
