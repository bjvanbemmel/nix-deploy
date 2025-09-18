package commands

import (
	_ "embed"
	"fmt"
)

//go:embed help_content.txt
var helpContent string

//go:embed config_content.txt
var configContent string

type Help struct{}

func (h Help) Execute(args ...string) error {
	if len(args) == 0 {
		fmt.Print(helpContent)
		return nil
	}

	switch args[0] {
	case "config":
		fmt.Print(configContent)
		return nil
	default:
		fmt.Print(helpContent)
		return nil
	}
}
