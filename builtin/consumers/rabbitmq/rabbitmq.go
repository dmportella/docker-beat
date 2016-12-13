package rabbitmq

import (
	"github.com/dmportella/docker-beat/plugin"
	_ "github.com/streadway/amqp" // Not currently used
)

type consumer struct {
}

func (consumer *consumer) OnEvent(events plugin.DockerEvent) {

}

func init() {
	// do something here
}
