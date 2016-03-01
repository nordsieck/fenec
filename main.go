package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nordsieck/fenec/build"
)

var (
	help = map[string]string{
		"": `Fenec is a tool for compiling Fenec source code.

Usage:

	fenec command [arguments]

The commands are:

	build	compile packages

Use "fenec help [command]" for more information about a command.
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
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			if err = build.ConvertDir(dir); err != nil {
				log.Fatal(err)
			}
		case "help":
			args = args[1:]
			if len(args) > 0 {
				txt, ok := help[args[0]]
				switch {
				case ok:
					fmt.Println(txt)
				default:
					fmt.Println(help[""])
				}
			} else {
				fmt.Println(help[""])
			}
		default:
			fmt.Println(help[""])
		}
	} else {
		fmt.Println(help[""])
	}
}
