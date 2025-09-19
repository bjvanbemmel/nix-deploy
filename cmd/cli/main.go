package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bjvanbemmel/go-templ/cmd/cli/commands"
	"github.com/bjvanbemmel/go-templ/config"
	"github.com/bjvanbemmel/go-templ/git"
)

func main() {
	unsanitizedArgs := os.Args[1:]
	var args []string
	var flags []string

	for _, arg := range unsanitizedArgs {
		if regexp.MustCompile("^(-[A-z-]+)").MatchString(arg) {
			flags = append(flags, arg)
			continue
		}

		args = append(args, arg)
	}

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

	git := git.New(conf)

	switch args[0] {
	case "config":
		configCmd := commands.NewConfig(conf)
		if err := configCmd.Execute(args[1:]...); err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		return
	case "fetch":
		fetch := commands.NewFetch(git)
		if err := fetch.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	case "version":
		commands.Version{}.Execute()
	case "help":
		fallthrough
	default:
		commands.Help{}.Execute(args[1:]...)
		return
	}
}
