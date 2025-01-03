/*
token-parser will parse and output claims from a JWT token

Usage:

    token-parser [flags] [path ...]

The commands are: 
    version - print the version of the tool 
    inspect - inspect a token (stdin) 

The flags are:

    -t string - the type of token being inspected (default "jwt")
    -v string -  the verifier type of token being inspected (default "none")
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bobmaertz/token-parser/internal/version"
	"github.com/bobmaertz/token-parser/pkg/inspector"
)

const (
	defaultType     = "jwt"
	defaultVerifier = "none"
)

var (
	// TODO: placeholders, not used yet
	tokenType string
	verifier  string

	commands = map[string]func([]string){
		"version": func(_ []string) {
			version.Print()
		},
		"inspect": func(s []string) {
			inspector := inspector.JwtNopInspector{}
			result, err := inspector.Inspect(s[0])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result.Claims)
		},
	}
)

func init() {
	flag.StringVar(&tokenType, "t", defaultType, "the type of token being inspected")
	flag.StringVar(&verifier, "v", defaultVerifier, "the verifier type of token being inspected")

	flag.Parse()
	flag.Usage = usage
}

func main() {
	args := flag.Args()

	if len(args) < 1 {
		usage()
		return
	}

	// Take the first argument as the command to run with the second argument as the argument to the command
	cmdName := args[0]
	f, ok := commands[cmdName]
	if !ok {
		fmt.Printf("command %v not recognized\n\n", cmdName)
		usage()
		return
	}

	f(args[1:])
}

func usage() {
	// TODO: Replace os.Args[0] with binary name.
	fmt.Fprintf(os.Stderr, "Usage of token-parser:\n")
	// TODO: Add description here.
	fmt.Fprintf(os.Stderr, "Command Usage:\n")
	fmt.Fprintf(os.Stderr, "  version       Print the version of the CLI\n")
	fmt.Fprintf(os.Stderr, "  inspect       Inspect a token\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Flag Usage:\n")

	flag.PrintDefaults()
}
