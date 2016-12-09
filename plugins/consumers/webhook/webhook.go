package webhook

import (
	"github.com/dmportella/docker-beat/plugins"
)

type consumer struct {
}

func (consumer *consumer) OnEvent(events plugins.DockerEvent) {

}

func init() {
	// do something here
}
