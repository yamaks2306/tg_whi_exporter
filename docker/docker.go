package docker

import (
	"context"
	"errors"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetPgContainerIP(containerName, networkName string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	var postgres *types.Container

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return "", err
	}

	for _, container := range containers {
		if strings.Contains(container.Names[0], containerName) {
			postgres = &container
		}
	}

	if postgres == nil {
		return "", errors.New("container postgres not found")
	}

	nw := postgres.NetworkSettings.Networks

	return nw[networkName].IPAddress, nil

}
