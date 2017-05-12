package main

import (
	"github.com/dmportella/docker-beat/logging"
	"github.com/dmportella/docker-beat/plugin"
	"github.com/fsouza/go-dockerclient"
)

func newDockerBeat(dockerEndpoint string, consumer string) (*dockerBeat, error) {
	dockerbeat := &dockerBeat{
		dockerEvents:   make(chan *docker.APIEvents),
		dockerEndpoint: dockerEndpoint,
		consumer:       consumer,
	}

	client, err := docker.NewClient(dockerbeat.dockerEndpoint)

	if err != nil {
		logging.Error.Printf(err.Error())
		return nil, err
	}

	dockerbeat.dockerClient = client

	return dockerbeat, nil
}

func (dockerbeat *dockerBeat) Start() {
	go dockerbeat.listContainers()

	go dockerbeat.dockerEventListener()
}

func (dockerbeat *dockerBeat) listContainers() {
	containers, _ := dockerbeat.dockerClient.ListContainers(docker.ListContainersOptions{All: true})

	for _, containerEntry := range containers {
		if container, _ := dockerbeat.dockerClient.InspectContainer(containerEntry.ID); container != nil {
			logging.Info.Printf("Container '%s' with ID '%s' %s.", container.Name, container.ID, container.State.Status)
		}
	}
}

func (dockerbeat *dockerBeat) dockerEventListener() {
	err := dockerbeat.dockerClient.AddEventListener(dockerbeat.dockerEvents)

	if err != nil {
		logging.Error.Printf(err.Error())
		panic(err)
	}

	for event := range dockerbeat.dockerEvents {
		logging.Info.Printf("Type: '%s' Action: '%s' Status: '%s' Time: '%d' Id: '%s'", event.Type, event.Action, event.Status, event.Time, event.ID)

		if dockerbeat.consumer != "console" {
			eventWrapper := plugin.DockerEvent{APIEvents: event}
			consumer := plugin.GetConsumer(dockerbeat.consumer)
			if consumer != nil {
				go consumer.OnEvent(eventWrapper)
			} else {
				logging.Error.Printf("Consumer '%s' is not available.", dockerbeat.consumer)
			}
		}
	}
}
