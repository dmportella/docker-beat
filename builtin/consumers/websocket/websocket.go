package rabbitmq

import (
	"github.com/dmportella/docker-beat/plugin"
	_ "golang.org/x/net/websocket"
)

type consumer struct {
}

func (consumer *consumer) OnEvent(event plugin.DockerEvent) {

}

func init() {
	// do something here
}

/*package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:8080/echo"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg)
}*/
