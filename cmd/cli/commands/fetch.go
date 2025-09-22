package commands

import (
	"fmt"

	"github.com/bjvanbemmel/go-templ/git"
)

type Fetch struct {
	Git git.GitIntegration
}

func NewFetch(git git.GitIntegration) Fetch {
	return Fetch{
		Git: git,
	}
}

func (f Fetch) Execute(args ...string) error {
	fmt.Println("Checking for new configurations...")

	err := f.Git.Pull()

	if err != nil {
		return err
	}

	rev, err := f.Git.LatestHash()
	if err != nil {
		return err
	}

	fmt.Printf("New configuration pulled with hash `%s`!\n", rev)

	return nil
}
