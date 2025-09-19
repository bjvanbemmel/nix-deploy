package commands

import (
	"errors"
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
	if errors.Is(err, git.ErrAlreadyUpToDate) {
		fmt.Println("No new configurations found!")
		return nil
	}

	if err != nil {
		return err
	}

	rev, err := f.Git.LatestHash()
	if err != nil {
		return err
	}
	fmt.Printf("New configuration pulled with hash `%s`!", rev)

	return nil
}
