package webhook

import (
	"bytes"
	"crypto/tls"
	"errors"
	"flag"
	"net/http"
	"net/url"
	"time"

	"github.com/dmportella/docker-beat/logging"
	"github.com/dmportella/docker-beat/plugin"
)

var (
	webHookEnpoint       string
	webhookSkipSSLVerify bool
)

const (
	defaultWebHookEndpoint = ""
	webHookEnpointUsage    = "webhook: The URL that events will be POSTed too."

	defaultSkipSSLVerify = false
	skipSSLVerifyUsage   = "webhook: Tells docker-beat to ignore ssl verification for the endpoint (not recommented)."

	userAgent = "Docker-Beat (https://github.com/dmportella/docker-beat, 0.0.0)"
)

type consumer struct {
	Debug bool
}

func (consumer *consumer) OnEvent(event plugin.DockerEvent, data []byte) {
	if _, err := url.Parse(webHookEnpoint); err != nil || webHookEnpoint == "" {
		logging.Error.Printf("Webhook url is not valid '%s'\n", webHookEnpoint)
	} else {
		consumer.request("POST", webHookEnpoint, data)
	}
}

func init() {
	flag.StringVar(&webHookEnpoint, "webhook-endpoint", defaultWebHookEndpoint, webHookEnpointUsage)
	flag.BoolVar(&webhookSkipSSLVerify, "webhook-skip-ssl-verify", defaultSkipSSLVerify, skipSSLVerifyUsage)

	consumer := &consumer{}

	plugin.RegisterConsumer("webhook", consumer)
}

func (consumer *consumer) request(method string, url string, b []byte) (response []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("user-agent", userAgent)

	tldConf := &tls.Config{
		InsecureSkipVerify: webhookSkipSSLVerify,
	}

	transport := &http.Transport{
		TLSClientConfig: tldConf,
	}

	httpClient := &http.Client{Timeout: (120 * time.Second), Transport: transport}

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
