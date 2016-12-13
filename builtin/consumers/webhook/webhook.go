package webhook

import (
	"flag"
	"github.com/dmportella/docker-beat/plugin"
)

var (
	webHookEnpoint string
)

const (
	defaultWebHookEndpoint = ""
	webHookEnpointUsage    = "The URL that events will be POSTed too."
)

type consumer struct {
}

func (consumer *consumer) OnEvent(events plugin.DockerEvent) {

}

func init() {
	// do something here
	flag.StringVar(&webHookEnpoint, "webhook-endpoint", defaultWebHookEndpoint, webHookEnpointUsage)
}
