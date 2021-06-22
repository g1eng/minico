package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (p *Provider) ContainerStop(conNameFilter string) error {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{Quiet: true})

	if err != nil {
		return err
	} else {
		//fuzzy name matching and stop containers recursively
		for _, container := range containers {
			name := strings.TrimLeft(container.Names[0], "/")
			if p.match(conNameFilter, name) {
				err = cli.ContainerStop(
					context.Background(),
					container.ID,
					nil,
				)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}
