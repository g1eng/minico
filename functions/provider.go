package functions

import (
	"fmt"
	"github.com/g1eng/minico/v2/functions/docker"
	"github.com/urfave/cli/v2"
)

// Unimplemented simply returns error for unimplemented features.
func Unimplemented(_ *cli.Context) error {
	return fmt.Errorf("This function is not implemented.\n")
}

// Provider is the interface for various container runtimes that
// is different in their APIs and designs.
type Provider interface {
	ContainerList(forAll bool, withIds bool, filter string) error
	ContainerRun(imageName string, autoRemove bool, command []string) error
	ImageList(withIds bool, filter string) error
	ContainerStop(filter string) error
}

// NewProvider creates a new instance of Provider implementation, and return it.
// For cosh, generated instance is used inside cli.Actionfunc.
func NewProvider(providerName string) Provider {
	switch providerName {
	case "docker":
		return &docker.Provider{}
	case "podman":
	case "lxc":
	case "k8s":
	default:
		panic(providerName + " not supported now!")
	}
	return nil
}
