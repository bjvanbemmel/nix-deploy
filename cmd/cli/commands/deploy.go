package commands

import (
	"fmt"

	"github.com/bjvanbemmel/go-templ/nix"
)

type Deploy struct {
	Nix nix.NixIntegration
}

func NewDeploy(nix nix.NixIntegration) Deploy {
	return Deploy{
		Nix: nix,
	}
}

func (d Deploy) Execute(args ...string) error {
	fmt.Println("Deploying configuration...")
	err := d.Nix.Switch()
	if err != nil {
		return err
	}

	fmt.Println("New configuration deployed!")

	return nil
}
