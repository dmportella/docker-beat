package main

import (
	"github.com/fsouza/go-dockerclient"
)

type dockerBeat struct {
	dockerEndpoint string
	consumer       string

	dockerEvents chan *docker.APIEvents
	dockerClient *docker.Client
}
