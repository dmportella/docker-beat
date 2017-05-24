package main

import (
	"github.com/fsouza/go-dockerclient"
)

type configuration struct {
	IndentJSON     bool
	DockerEndpoint string
	Consumer       string
}

type dockerBeat struct {
	dockerEvents chan *docker.APIEvents
	dockerClient *docker.Client
	config       configuration
}
