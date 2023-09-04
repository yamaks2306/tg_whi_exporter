package docker

import (
	"context"
	"fmt"

	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/yamaks2306/tg_whi_exporter/util"
)

func GetPgContainerIP(containerName, networkName string) string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	util.CheckError(err)

	ctx := context.Background()

	var postgres types.Container

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	util.CheckError(err)

	for _, container := range containers {
		if strings.Contains(container.Names[0], containerName) {
			postgres = container
		}
	}

	nw := postgres.NetworkSettings.Networks
	fmt.Println(nw[networkName].IPAddress)

	return nw[networkName].IPAddress

}
