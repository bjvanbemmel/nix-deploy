package commands

import (
	"encoding/json"
	"fmt"

	"github.com/bjvanbemmel/go-templ/config"
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
		c.Set(args[1], args[2])
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
	switch variable {
	case "source":
		c.Configuration.Source = value
	case "flake":
		c.Configuration.Flake = value
	}

	return c.Configuration.Save()
}
