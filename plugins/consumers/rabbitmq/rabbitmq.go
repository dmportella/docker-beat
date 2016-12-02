package rabbitmq

import (
	"github.com/dmportella/docker-beat/plugins"
	_ "github.com/streadway/amqp" // Not currently used
)

type consumer struct {
}

func (consumer *consumer) OnEvent(events plugins.DockerEvent) {

}

func init() {
	// do something here
}