package docker

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ContainerListFunc is the function type which is used in recursive actions
// over queried container list.
type ContainerListFunc func(forAll bool, forStopped bool, nameFilter string) (err error)

// ContainerList lists containers with specified filters.
func (p *Provider) ContainerList(forAll bool, withIds bool, nameFilter string) (err error) {
	var (
		cli        *client.Client
		containers []types.Container
	)
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	containers, err = cli.ContainerList(p.getConListCondition(forAll))
	if err != nil {
		return err
	}

	return p.queryConByName(containers, withIds, nameFilter)
}

func (p *Provider) getConListCondition(forAll bool) (context.Context, types.ContainerListOptions) {
	listOpts := types.ContainerListOptions{Quiet: false}
	if forAll {
		listOpts.All = true
	}
	return context.Background(), listOpts
}

func (p *Provider) queryConByName(containers []types.Container, withIds bool, nameFilter string) error {
	found := false
	//list and filter container with names
	for _, container := range containers {
		name := strings.TrimLeft(container.Names[0], "/")
		// filter container with nameFilter
		if p.match(nameFilter, name) {
			if withIds {
				fmt.Printf("%s %s\n", strings.TrimLeft(container.ID, "sha256:")[:10], name)
			} else {
				fmt.Printf("%s\n", name)
			}
			found = true
		}
	}
	if !found {
		return errors.New("")
	}
	return nil
}
