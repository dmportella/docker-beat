# docker-beat
A simple docker event beat server that will distribute docker events to plugins/actors to perform actions against them.

@dmportella

[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/dmportella/docker-beat/master/LICENSE) [![Build Status](https://travis-ci.org/dmportella/docker-beat.svg?branch=master)](https://travis-ci.org/dmportella/docker-beat) [![GoDoc](https://godoc.org/github.com/dmportella/docker-beat?status.svg)](https://godoc.org/github.com/dmportella/docker-beat) [![Go Report Card](https://goreportcard.com/badge/github.com/dmportella/docker-beat)](https://goreportcard.com/report/github.com/dmportella/docker-beat) [![Github Release](https://img.shields.io/github/release/dmportella/docker-beat.svg)](https://github.com/dmportella/docker-beat/releases)

## Dockerhub

[![dockeri.co](http://dockeri.co/image/dmportella/docker-beat)](https://hub.docker.com/r/dmportella/docker-beat/)

## Current support

Currently `docker-beat` supports sending docker events to the following endpoints.

 - console
 - rabbitmq
 - webhook
 - websockets

See projects for future endpoints support and additional things on the roadmap.

## Running in Docker

The docker container supports Docker API Socket as a volume (not recommended) or you can provide the Docker API Url.

### Running docker-beat with Docker API as Socker Volume

> $ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock dmportella/docker-beat --consumer webhook --webhook-endpoint http://requestb.in/rn7cixrn

### Running docker-beat with Docker API as Endpoint

> $ docker run --rm dmportella/docker-beat --consumer webhook --docker-endpoint "tcp://localhost:2375" --webhook-endpoint http://requestb.in/rn7cixrn

## Running locally

Should be simple to run the application locally it differs just slightly between OS.

```
docker-beat - Version: 0.0.5 Branch: master Revision: 55ea00b. OSArch: linux/amd64.
Daniel Portella (c) 2016
Usage of ./bin/docker-beat:
  -consumer string
      Consumer to use: Webhook, Rabbitmq, etc. (default "console")
  -docker-endpoint string
      The Url or unix socket address for the Docker Remote API. (default "unix:///var/run/docker.sock")
  -help
      Prints the help information.
  -rabbitmq-endpoint string
      rabbitmq: The URL that events will be published too. (default "amqp://guest:guest@localhost:5672/")
  -rabbitmq-exchange string
      rabbitmq: The exchange docker-beat will publish messages too. (default "docker-beat")
  -rabbitmq-exchange-type string
      rabbitmq: The exchange type that docker-beat will create/connect too. (direct|fanout|topic|x-custom) (default "fanout")
  -rabbitmq-reliable
      rabbitmq: The ensures messages published are confirmed.
  -rabbitmq-routing-key string
      rabbitmq: The routing key for messages published to the exchange. (default: docker-event (default "docker-event")
  -verbose
      Redirect trace information to the standard out.
  -webhook-endpoint string
      webhook: The URL that events will be POSTed too.
  -webhook-indent
      webhook: Indent the json output.
  -webhook-skip-ssl-verify
      webhook: Tells docker-beat to ignore ssl verification for the endpoint (not recommented).
  -websocket-endpoint string
      websocket: The URL that events will be streamed too.
  -websocket-origin string
      websocket: The origin of the request to be used in the web socket stream.
  -websocket-protocol string
      websocket: The protocol to be used in the web socket stream.
```

### Linux, Darwin and FreeBSD

> ./docker-beat --consumer webhook --webhook-endpoint http://requestb.in/rn7cixrn

### Windows

> .\docker-beat --consumer webhook --docker-endpoint "tcp://localhost:2375" --webhook-endpoint http://requestb.in/rn7cixrn
