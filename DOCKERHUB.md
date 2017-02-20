A simple docker event beat server that will distribute docker events to plugins/actors to perform actions against them.

# Supported tags and respective `Dockerfile` links

* [`0.0.2`, (scratch/dockerfile)](https://github.com/dmportella/docker-beat/blob/0.0.2/dockerfile)
* [`0.0.3`, (scratch/dockerfile)](https://github.com/dmportella/docker-beat/blob/0.0.3/dockerfile)
* [`0.0.5`, `latest` (scratch/dockerfile)](https://github.com/dmportella/docker-beat/blob/0.0.5/dockerfile)

## Running in Docker

The docker container supports Docker API Socket as a volume (not recommended) or you can provide the Docker API Url.

### Running docker-beat with Docker API as Socker Volume

> $ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock dmportella/docker-beat --consumer webhook --webhook-endpoint http://requestb.in/rn7cixrn

### Running docker-beat with Docker API as Endpoint

> $ docker run --rm dmportella/docker-beat --consumer webhook --docker-endpoint "tcp://localhost:2375" --webhook-endpoint http://requestb.in/rn7cixrn

### Commands

See below the list of available arguments.

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
