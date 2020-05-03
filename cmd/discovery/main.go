package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"

	"encoding/json"
)

var containerName string = "docker-discovery"

func main() {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var networks []string

	for _, container := range containers {
		for _, name := range container.Names {
			if name == containerName {
				//self = container

				networks := make([]string, 0, len(container.NetworkSettings.Networks))
				for k := range container.NetworkSettings.Networks {
					networks = append(networks, k)
				}
			}
		}
	}

	if len(networks) <= 1 {
		networkList, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
		if err != nil {
			panic(err)
		}
		for _, network := range networkList {
			networks = append(networks, network.Name)
		}
	}

	type target struct {
		Targets []string          `jsonOutput:"targets"`
		Labels  map[string]string `jsonOutput:"labels"`
	}

	var targets []target
	for _, container := range containers {
		var host string

		for _, network := range networks {
			if val, ok := container.NetworkSettings.Networks[network]; ok {
				host = val.IPAddress
				break
			}
		}

		if host == "" {
			break
		}

		var str strings.Builder
		var m map[string]string

		m = make(map[string]string)
		m["job"] = container.Names[0]

		str.WriteString(host)
		str.WriteString(":")
		str.WriteString("9090")
		//str.WriteString("/metrics")

		targets = append(targets, target{
			Targets: []string{0: str.String()},
			Labels:  m,
		})
	}

	jsonOutput, _ := json.Marshal(targets)
	fmt.Println(string(jsonOutput))

	if string(jsonOutput) == "null" {
		os.Exit(0)
	}

	err = ioutil.WriteFile("/etc/prometheus/targets/docker-targets.jsonOutput", jsonOutput, 0644)
	if err != nil {
		panic(err)
	}
}
