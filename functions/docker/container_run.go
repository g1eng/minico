package docker

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

// ExecResult represents a result returned from Exec()
type ExecResult struct {
	ExitCode  int
	outBuffer *bytes.Buffer
	errBuffer *bytes.Buffer
}

func (p *Provider) ContainerRun(imageName string, persistent bool, command []string) (err error) {
	var (
		cli *client.Client
		con container.ContainerCreateCreatedBody
	)
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	//name auto configuration
	containerName := regexp.
		MustCompile("(^[^/]+/|:[^:]+$)").
		ReplaceAllString(imageName, "")
	if !persistent {
		containerName += "_tmp"
	}

	ctx := context.Background()
	con, err = cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: imageName,
			Cmd:   command,
		},
		&container.HostConfig{
			AutoRemove: !persistent,
		},
		&network.NetworkingConfig{},
		&specs.Platform{
			Architecture: "amd64",
			OS:           "linux",
		},
		containerName,
	)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(
		ctx,
		con.ID,
		types.ContainerStartOptions{},
	)
	if err != nil {
		return err
	}

	result, err := Exec(ctx, cli, con.ID, command)
	if err != nil {
		return err
	}

	if result.ExitCode == 0 {
		_, _ = io.Copy(os.Stdout, result.outBuffer)
	} else if result.ExitCode > 0 {
		_, _ = io.Copy(os.Stderr, result.outBuffer)
	} else {
		return errors.New("unknown exit code")
	}
	if !persistent {
		return p.ContainerStop(con.ID)
	}

	return nil

}

func Exec(ctx context.Context, cli client.APIClient, id string, cmd []string) (ExecResult, error) {
	// prepare exec
	execConfig := types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Detach:       false,
		Cmd:          cmd,
	}
	cresp, err := cli.ContainerExecCreate(ctx, id, execConfig)
	if err != nil {
		return ExecResult{}, err
	}
	execID := cresp.ID

	// run it, with stdout/stderr attached
	aresp, err := cli.ContainerExecAttach(ctx, execID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}
	defer aresp.Close()

	// read the output
	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error, 1)

	go func() {
		// StdCopy demultiplexes the stream into two buffers
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, aresp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return ExecResult{}, err
		}
		break

	case <-ctx.Done():
		return ExecResult{}, ctx.Err()
	}

	// get the exit code
	iresp, err := cli.ContainerExecInspect(ctx, execID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{ExitCode: iresp.ExitCode, outBuffer: &outBuf, errBuffer: &errBuf}, nil
}
