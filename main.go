package main

import (
	"bufio"
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"golang.org/x/crypto/blake2b"
)

const version = "0.0.1"

const usage = "" +
	`keyx computes funny hash of input data.

Usage:
  keyx -c <case>| --case=<case>
  keyx -v | --version
  keyx -h | --help

Options:
  -c <case>, --case=<case>  Case used in hash string: 0 is lower, 1 is upper.
  -v, --version             Show version.
  -h, --help                Show this screen.
`

func main() {
	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad arguments")
		os.Exit(1)
	}
	if arguments["--version"].(bool) {
		fmt.Println(version)
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	data, _ := reader.ReadBytes('\n')
	hash := blake2b.Sum256(data)
	switch c := arguments["--case"]; {
	case c == "0":
		fmt.Printf("%x\n", hash)
	case c == "1":
		fmt.Printf("%X\n", hash)
	default:
		fmt.Fprintln(os.Stderr, "bad case indicator")
		os.Exit(2)
	}
}
