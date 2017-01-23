# docker-beat
A simple docker event beat server that will distribute docker events to plugins/actors to perform actions against them.

@dmportella

[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/dmportella/docker-beat/master/LICENSE) [![Build Status](https://travis-ci.org/dmportella/docker-beat.svg?branch=master)](https://travis-ci.org/dmportella/docker-beat) [![GoDoc](https://godoc.org/github.com/dmportella/docker-beat?status.svg)](https://godoc.org/github.com/dmportella/docker-beat) [![Go Report Card](https://goreportcard.com/badge/github.com/dmportella/docker-beat)](https://goreportcard.com/report/github.com/dmportella/docker-beat) [![Github Release](https://img.shields.io/github/release/dmportella/docker-beat.svg)](https://github.com/dmportella/docker-beat/releases)

## Dockerhub

[![dockeri.co](http://dockeri.co/image/dmportella/docker-beat)](https://hub.docker.com/r/dmportella/docker-beat/)

## Documentation

Please check the wiki for more information: [WIKI](https://github.com/dmportella/docker-beat/wiki)

## Running in Docker

The docker container supports Docker API Socket as a volume (not recommended) or you can provide the Docker API Url (current does not support SSL).

### Running docker-beat with Docker API as Socker Volume

> $ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock dmportella/docker-beat --consumer webhook --webhook-endpoint http://requestb.in/rn7cixrn

### Running docker-beat with Docker API as Endpoint

> $ docker run --rm dmportella/docker-beat --consumer webhook --docker-endpoint "tcp://localhost:2375" --webhook-endpoint http://requestb.in/rn7cixrn