package main

import (
	"fmt"
	"os"

	"github.com/bjvanbemmel/go-templ/cmd/cli/commands"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		commands.Help{}.Execute()
		return
	}

	switch args[0] {
	case "help":
		commands.Help{}.Execute()
		return
	}

	fmt.Println("Hello, World!")
}
