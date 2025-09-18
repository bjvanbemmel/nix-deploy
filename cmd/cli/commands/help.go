package commands

import (
	_ "embed"
	"fmt"
)

//go:embed help_content.txt
var content string

type Help struct{}

func (h Help) Execute(args ...string) error {
	fmt.Println(content)
	return nil
}
