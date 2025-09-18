package main

import (
	"fmt"
	"os"

	"github.com/bjvanbemmel/go-templ/cmd/cli/commands"
	"github.com/bjvanbemmel/go-templ/config"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		commands.Help{}.Execute()
		return
	}

	conf := config.New()
	if err := conf.Load(); err != nil {
		if err := conf.Save(); err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	}

	switch args[0] {
	case "config":
		configCmd := commands.NewConfig(conf)
		if err := configCmd.Execute(args[1:]...); err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		return
	case "version":
		commands.Version{}.Execute()
	case "help":
		fallthrough
	default:
		commands.Help{}.Execute(args[1:]...)
		return
	}
}
