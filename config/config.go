package config

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
)

const (
	DEFAULT_SOURCE             string = "https://github.com/bjvanbemmel/nix-config"
	DEFAULT_FLAKE              string = "~/.config/nix"
	GLOBAL_CONFIG_DEFAULT_PATH string = "~/.config/nix-deploy/config"
)

type Configuration struct {
	Source string `json:"source"`
	Flake  string `json:"flake"`
}

func New() Configuration {
	return Configuration{
		Source: DEFAULT_SOURCE,
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
	// Replace `~` with the value of ${HOME}
	return regexp.MustCompile("^~").ReplaceAllString(GLOBAL_CONFIG_DEFAULT_PATH, os.Getenv("HOME"))
}
