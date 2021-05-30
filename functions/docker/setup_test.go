package docker

import (
	"github.com/docker/docker/api/types"
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	Suite(&testSuite{
		provider: Provider{
			Hostname:    "",
			environment: nil,
		},
		fixtures: testFixture{
			images: []types.ImageSummary{
				{
					Containers: 123,
					Created:    123,
					ID:         "123456712345671234567123456712345671234567",
					Labels:     map[string]string{"this": "dev"},
					Size:       123,
					SharedSize: 12,
					RepoTags:   []string{"dev"},
				},
			},
			con: []types.Container{
				{
					ID:         "123456123456123456123456123456123456124356",
					Names:      []string{"ok"},
					Image:      "dev",
					ImageID:    "123456712345671234567123456712345671234567",
					Command:    "ls",
					Created:    1234567891234567890,
					SizeRootFs: 12345,
					SizeRw:     123,
					Labels:     map[string]string{"this": "okd"},
					State:      "ok",
					Status:     "success",
				},
			},
		},
	})
}

type testSuite struct {
	provider Provider
	fixtures testFixture
}
type testFixture struct {
	images          []types.ImageSummary
	con             []types.Container
	conStartOpts    []types.ContainerStartOptions
	conAttachOpts   []types.ContainerAttachOptions
	conState        []types.ContainerState
	conExecInspects []types.ContainerExecInspect
}

func Test(t *testing.T) { TestingT(t) }
