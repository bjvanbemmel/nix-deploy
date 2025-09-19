package config

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
)

const (
	DEFAULT_SOURCE             string = "git@github.com:bjvanbemmel/nix-config"
	DEFAULT_BRANCH             string = "main"
	DEFAULT_REMOTE             string = "origin"
	DEFAULT_FLAKE              string = "~/.config/nix"
	GLOBAL_CONFIG_DEFAULT_PATH string = "~/.config/nix-deploy/config"
)

type Configuration struct {
	Source string `json:"source"`
	Branch string `json:"branch"`
	Remote string `json:"remote"`
	Flake  string `json:"flake"`
	Hash   string `json:"hash"`
}

func New() Configuration {
	return Configuration{
		Source: DEFAULT_SOURCE,
		Branch: DEFAULT_BRANCH,
		Remote: DEFAULT_REMOTE,
		Flake:  DEFAULT_FLAKE,
	}
}

func (c *Configuration) Load() error {
	if _, err := os.Stat(c.Path()); err != nil {
		return err
	}

	raw, err := os.ReadFile(c.Path())
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, c); err != nil {
		return err
	}

	return nil
}

func (c Configuration) Save() error {
	raw, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path.Dir(c.Path())); err != nil {
		if err := os.MkdirAll(path.Dir(c.Path()), os.ModePerm); err != nil {
			return err
		}
	}

	if err := os.WriteFile(c.Path(), raw, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (c Configuration) Path() string {
	return c.toAbsolute(GLOBAL_CONFIG_DEFAULT_PATH)
}

func (c Configuration) AbsoluteSource() string {
	return c.toAbsolute(c.Source)
}

func (c Configuration) AbsoluteFlake() string {
	return c.toAbsolute(c.Flake)
}

func (c Configuration) toAbsolute(path string) string {
	// Replace `~` with the value of ${HOME}
	return regexp.MustCompile("^~").ReplaceAllString(path, os.Getenv("HOME"))
}
