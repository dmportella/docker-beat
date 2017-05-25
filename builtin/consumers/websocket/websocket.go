package rabbitmq

import (
	"flag"

	"github.com/dmportella/docker-beat/logging"
	"github.com/dmportella/docker-beat/plugin"
	"golang.org/x/net/websocket"
)

var (
	// WebsocketOrigin The http origin for the web socket.
	WebsocketOrigin string

	// WebsocketProtocol The protocol for the web socket.
	WebsocketProtocol string

	// WebsocketEndpoint The endpoint for the web socket.
	WebsocketEndpoint string
)

const (
	defaultWebsocketEndpoint = ""
	websocketEndpointUsage   = "websocket: The URL that events will be streamed too."

	defaultWebsocketProtocol = ""
	websocketProtocolUsage   = "websocket: The protocol to be used in the web socket stream."

	defaultWebsocketOrigin = ""
	websocketOriginUsage   = "websocket: The origin of the request to be used in the web socket stream."

	userAgent = "Docker-Beat (https://github.com/dmportella/docker-beat, 0.0.0)"
)

type consumer struct {
	socket *websocket.Conn
}

func (consumer *consumer) resetConnection() {
	consumer.socket.Close()
	consumer.socket = nil
}

func (consumer *consumer) OnEvent(event plugin.DockerEvent, data []byte) {
CONN:
	count := 0
	if consumer.socket == nil {
		ws, err := websocket.Dial(WebsocketEndpoint, WebsocketProtocol, WebsocketOrigin)
		if err != nil {
			logging.Error.Println(err.Error())
			return
		}

		consumer.socket = ws
	}

	_, err := consumer.socket.Write(data)
	if err != nil {
		logging.Error.Printf(err.Error())
		consumer.resetConnection()
		if count < 2 {
			count = count + 1
			goto CONN
		}
	}
}

func init() {
	flag.StringVar(&WebsocketEndpoint, "websocket-endpoint", defaultWebsocketEndpoint, websocketEndpointUsage)
	flag.StringVar(&WebsocketProtocol, "websocket-protocol", defaultWebsocketProtocol, websocketProtocolUsage)
	flag.StringVar(&WebsocketOrigin, "websocket-origin", defaultWebsocketOrigin, websocketOriginUsage)

	consumer := &consumer{}

	plugin.RegisterConsumer("websocket", consumer)
}
