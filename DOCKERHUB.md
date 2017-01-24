A simple docker event beat server that will distribute docker events to plugins/actors to perform actions against them.

# Supported tags and respective `Dockerfile` links

* [`0.0.1`, `latest` (scratch/dockerfile)](https://github.com/dmportella/docker-beat/blob/0.0.1/dockerfile)

## Running in Docker

The docker container supports Docker API Socket as a volume (not recommended) or you can provide the Docker API Url.

### Running docker-beat with Docker API as Socker Volume

> $ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock dmportella/docker-beat --consumer webhook --webhook-endpoint http://requestb.in/rn7cixrn

### Running docker-beat with Docker API as Endpoint

> $ docker run --rm dmportella/docker-beat --consumer webhook --docker-endpoint "tcp://localhost:2375" --webhook-endpoint http://requestb.in/rn7cixrn
