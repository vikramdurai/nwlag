/*
	Nwlag is an estoric programming language that compiles (nwlag_c)
	into human-readable intermediate code (JSON), which is executable
	using nwlag_i.

	Feedback is welcome.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var help_msg string = `Usage: nwlag [command] [arguments]
Possible commands:
	compile: compile a Nwlag source file
	run: run output from nwlag compile
	help: print this help message

Current version: 0.1dev
Feedback is welcome at https://github.com/vikramdurai/nwlag/issues`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(help_msg)
		os.Exit(0)
	}
	// find out which subcommand
	switch os.Args[1] {
	case "run":
		// find argument to subcommand
		v := os.Args[2]
		err := Execute(v)
		if err != nil {
			fmt.Printf("Error while executing: %s\n", err)
		}
	case "compile":
		// find argument to subcommand
		v := os.Args[2]
		// open given file
		code, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Printf("Error while processing: %s\n", err)
		}
		// Compile's signature takes in
		// string, not []byte
		out, err := Compile(string(code))
		if err != nil {
			fmt.Printf("Error while compiling: %s\n", err)
		}
		// change the extension
		fname := strings.TrimSuffix(v, ".nwl")
		fname += ".json"
		// write compiled output to the same file
		// with file permission mode 777
		err = ioutil.WriteFile(fname, out, 0666)
		if err != nil {
			fmt.Printf("Error while writing output: %s\n", err)
		}
	case "help":
		fmt.Println(help_msg)

	default:
		// It doesn't hurt folks to be told
		// again, how it works
		fmt.Println(help_msg)
	}
}
