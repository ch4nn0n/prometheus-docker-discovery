package main

import (
	"context"
	"fmt"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func main() {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	for _, network := range networks {
		fmt.Println(network.Name)
	}

	for _, container := range containers {
		//fmt.Printf("%s %s\n", container.ID[:10], container.Image, container.State)

		fmt.Println(container.State == "running")
		fmt.Println(container.Labels)
		fmt.Println(container.NetworkSettings)
		fmt.Println(container.Ports[0].PublicPort)
	}
}
