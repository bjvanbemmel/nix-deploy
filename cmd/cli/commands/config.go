package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bjvanbemmel/go-templ/config"
)

var (
	ErrInvalidVariable error = errors.New("invalid variable given")
)

type Config struct {
	Configuration config.Configuration
}

func NewConfig(config config.Configuration) Config {
	return Config{
		Configuration: config,
	}
}

func (c *Config) Execute(args ...string) error {
	if len(args) == 0 {
		Help{}.Execute("config")
		return nil
	}

	switch args[0] {
	case "list":
		c.List()
	case "set":
		if len(args[1:]) < 2 {
			return ErrInvalidArgument
		}
		return c.Set(args[1], args[2])
	case "path":
		fmt.Println(c.Configuration.Path())
	}

	return nil
}

func (c Config) List() error {
	raw, err := json.MarshalIndent(c.Configuration, "", "	")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", raw)

	return nil
}

func (c *Config) Set(variable, value string) error {
	var old string

	switch variable {
	case "source":
		old = c.Configuration.Source
		c.Configuration.Source = value
	case "branch":
		old = c.Configuration.Branch
		c.Configuration.Branch = value
	case "remote":
		old = c.Configuration.Remote
		c.Configuration.Remote = value
	case "flake":
		old = c.Configuration.Flake
		c.Configuration.Flake = value
	case "hash":
		old = c.Configuration.Flake
		c.Configuration.Hash = value
	default:
		return ErrInvalidVariable
	}

	if err := c.Configuration.Save(); err != nil {
		return err
	}

	fmt.Printf("Variable `%s` has been updated from `%s` to `%s`\n", variable, old, value)

	return nil
}
