package git

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/bjvanbemmel/go-templ/config"
)

const (
	PULL_NO_RESULTS string = "Already up to date."
)

var (
	ErrAlreadyUpToDate error = errors.New("already up to date")
)

type GitIntegration struct {
	Config config.Configuration
}

func New(config config.Configuration) GitIntegration {
	return GitIntegration{
		Config: config,
	}
}

func (g *GitIntegration) Pull() error {
	oldHash, err := g.LatestHash()
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = g.Config.AbsoluteFlake()

	raw, err := cmd.Output()
	if err != nil {
		return err
	}

	rev, err := g.LatestHash()
	if err != nil {
		return err
	}

	if oldHash == rev {
		return ErrAlreadyUpToDate
	}

	g.Config.Hash = rev
	g.Config.Save()

	if strings.TrimSuffix(string(raw), "\n") == PULL_NO_RESULTS {
		return ErrAlreadyUpToDate
	}

	return nil
}

func (g GitIntegration) LatestHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = g.Config.AbsoluteFlake()

	raw, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(raw), "\n"), nil
}
