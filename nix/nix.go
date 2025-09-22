package nix

import (
	"os/exec"

	"github.com/bjvanbemmel/go-templ/config"
)

type NixIntegration struct {
	Config config.Configuration
}

func New(config config.Configuration) NixIntegration {
	return NixIntegration{
		Config: config,
	}
}

func (n NixIntegration) Switch() error {
	cmd := exec.Command("nixos-rebuild", "switch", "--flake", n.Config.AbsoluteFlake())

	return cmd.Run()
}
