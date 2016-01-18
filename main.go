package main

import (
	"fmt"
	"os"
)

var (
	help = map[string]string{
		"": `Wendigo is a tool for compiling Wendigo source code.

Usage:

	wendigo command [arguments]

The commands are:

	build	compile packages

Use "wendigo help [command]" for more information about a command.
`,
		"build": `Build compiles the current package and all sub-packages
recursively.  It has no arguments.`,
	}
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "build":
			fmt.Println("Building!")
		case "help":
			fmt.Println("Helping!")
		default:
			fmt.Println(help[""])
		}
		return
	}
	fmt.Println(help[""])
}
