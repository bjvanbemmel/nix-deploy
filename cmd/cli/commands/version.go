package commands

import "fmt"

type Version struct{}

const (
	CURRENT_VERSION string = "v0.1.0-dev"
)

func (v Version) Execute(args ...string) error {
	fmt.Println(CURRENT_VERSION)
	return nil
}
