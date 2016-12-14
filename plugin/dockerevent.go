package plugin

import (
	"github.com/fsouza/go-dockerclient"
)

// DockerEvent encapsulates a docker event please refer to 'http://github.com/fsouza/go-dockerclient/event.go' for more information.
type DockerEvent struct {
	*docker.APIEvents
}
