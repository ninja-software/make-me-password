package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	mmp "github.com/ninja-software/make-me-password"
)

// Version of the software
const Version string = "v0.1.2"

func main() {
	format := flag.String("format", "base64", "output format. choices: base64, hex")
	ver := flag.Bool("version", false, "version of the software")
	flag.Parse()

	if *ver {
		fmt.Printf("%s version %s\n", os.Args[0], Version)
		os.Exit(0)
	}

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
