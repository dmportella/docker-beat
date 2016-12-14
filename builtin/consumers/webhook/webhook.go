package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"github.com/dmportella/docker-beat/logging"
	"github.com/dmportella/docker-beat/plugin"
	"net/http"
	"time"
)

var (
	webHookEnpoint string
)

const (
	defaultWebHookEndpoint = ""
	webHookEnpointUsage    = "webhook: The URL that events will be POSTed too."

	userAgent = "Docker-Beat (https://github.com/dmportella/docker-beat, 0.0.0)"
)

type consumer struct {
	Debug bool
}

func (consumer *consumer) OnEvent(event plugin.DockerEvent) {
	data, _ := json.MarshalIndent(event, "", "    ")

	consumer.request("POST", "http://requestb.in/vheq9vvh", data)
}

func init() {
	// do something here
	flag.StringVar(&webHookEnpoint, "webhook-endpoint", defaultWebHookEndpoint, webHookEnpointUsage)

	consumer := &consumer{}

	plugin.RegisterConsumer(consumer)
}

func (consumer *consumer) request(method string, url string, b []byte) (response []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("user-agent", userAgent)

	httpClient := &http.Client{Timeout: (120 * time.Second)}

	res, err := httpClient.Do(req)

	if err != nil {
		logging.Warning.Println("Request error", err)
		err = errors.New("Http request returned an error")
		return
	}

	defer res.Body.Close()

	if consumer.Debug {
		logging.Trace.Printf("API REQUEST\tURL :: %s\n", url)
		logging.Trace.Printf("API RESPONSE\tSTATUS :: %s\n", res.Status)
		for k, v := range res.Header {
			logging.Trace.Printf("API RESPONSE\tHEADER :: [%s] = %+v\n", k, v)
		}
		logging.Trace.Printf("API RESPONSE\tBODY :: [%s]\n", response)
	}
	return
}
