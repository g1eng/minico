package docker

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

func (p *Provider) ImageList(withIds bool, nameFilter string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(
		context.Background(),
		types.ImageListOptions{},
	)

	if err != nil {
		panic(err)
	}

	return p.queryImageByName(images, withIds, nameFilter)
}

func (p *Provider) queryImageByName(images []types.ImageSummary, withIds bool, nameFilter string) error {
	found := false
	if withIds != true {
		for _, image := range images {
			for _, tag := range image.RepoTags {
				if p.match(nameFilter, tag) {
					fmt.Printf("%s\n", tag)
					found = true
				}
			}
		}
	} else {
		for _, image := range images {
			for _, tag := range image.RepoTags {
				if p.match(nameFilter, tag) {
					fmt.Printf("%s %s\n", strings.Trim(image.ID, "sha256:")[0:10], tag)
					found = true
				}
			}
		}
	}
	if found == false {
		return errors.New("") //image not found
	}
	return nil
}
