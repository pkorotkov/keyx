package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"golang.org/x/crypto/blake2b"
)

const version = "0.0.1"

const usage = "" +
	`keyx computes a hash of input data.

Usage:
  keyx [-n | --no-prefix] [-u | --upper-case] [--hash=<hash>]
  keyx -v | --version
  keyx -h | --help

Options:
  -n, --no-prefix   Do not use prefix before data read from stdin. 
  -u, --upper-case  Use upper case (instead of lower case as default) for hex string.
  --hash=<hash>     Set the hash algorithm (blake2b-256, sha-1, sha-256) [default: blake2b-256].
  -v, --version     Show version.
  -h, --help        Show this screen.
`

func main() {
	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: bad arguments")
		os.Exit(1)
	}
	if arguments["--version"].(bool) {
		fmt.Println(version)
		os.Exit(0)
	}
	if arguments["--no-prefix"].(bool) {
		fmt.Print("")
	} else {
		fmt.Print("> ")
	}
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadBytes('\n')
	data = bytes.TrimSpace(data)
	var hash []byte
	switch arguments["--hash"] {
	case "blake2b-256":
		h := blake2b.Sum256(data)
		hash = h[:]
	case "sha-1":
		h := sha1.Sum(data)
		hash = h[:]
	case "sha-256":
		h := sha256.Sum256(data)
		hash = h[:]
	default:
		fmt.Fprintln(os.Stderr, "error: unknown hash algorithm")
		os.Exit(2)
	}
	if !arguments["--upper-case"].(bool) {
		fmt.Printf("%x\n", hash)
	} else {
		fmt.Printf("%X\n", hash)
	}
}
