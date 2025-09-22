package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/bjvanbemmel/go-templ/cmd/cli/commands"
	"github.com/bjvanbemmel/go-templ/config"
	"github.com/bjvanbemmel/go-templ/git"
	"github.com/bjvanbemmel/go-templ/nix"
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

	gitIntegration := git.New(conf)
	nixIntegration := nix.New(conf)

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
		fetch := commands.NewFetch(gitIntegration)
		if err := fetch.Execute(); errors.Is(err, git.ErrAlreadyUpToDate) {
			fmt.Println("No new configurations found!")
			return
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	case "deploy":
		fetch := commands.NewFetch(gitIntegration)
		if err := fetch.Execute(); errors.Is(err, git.ErrAlreadyUpToDate) {
			fmt.Println("No new configurations found!")
			return
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		deploy := commands.NewDeploy(nixIntegration)
		if err := deploy.Execute(); err != nil {
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
