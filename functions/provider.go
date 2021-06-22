package functions

import "github.com/g1eng/minico/functions/docker"

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
