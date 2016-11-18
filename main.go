package main

import (
	"flag"
	"fmt"
	"github.com/dmportella/docker-beat/logging"
	"github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"os"
	"os/signal"
)

// Set on build
var (
	Build    string
	Branch   string
	Revision string
	OSArch   string
)

// Variables used for command line parameters
var (
	DockerEndpoint string
	Version        bool
	Verbose        bool
)

func init() {
	const (
		defaultDockerEndpoint = "unix:///var/run/docker.sock"
		dockerEndpointUsage   = "The Url or unix socket address for the Docker Remote API."
	)

	flag.StringVar(&DockerEndpoint, "docker-endpoint", defaultDockerEndpoint, dockerEndpointUsage)

	const (
		defaultVerbose = false
		verboseUsage   = "Redirect trace information to the standard out."
	)

	flag.BoolVar(&Verbose, "verbose", defaultVerbose, verboseUsage)
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	fmt.Printf("docker-beat - Version: %s Branch: %s Revision: %s. OSArch: %s.\n\rDaniel Portella (c) 2016\n\r", Build, Branch, Revision, OSArch)

	if Verbose {
		logging.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	} else {
		logging.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				logging.Info.Printf("Application ended on %s\n\r", sig)

				os.Exit(0)
			}
		}
	}()

	dockerEvents := make(chan *docker.APIEvents)

	client, err := docker.NewClient(DockerEndpoint)

	if err != nil {
		logging.Error.Printf(err.Error())
	}

	go listContainers(client)

	go dockerEventListener(dockerEvents, client)

	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}

func listContainers(client *docker.Client) {
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})

	for _, containerEntry := range containers {
		if container, _ := client.InspectContainer(containerEntry.ID); container != nil {
			logging.Info.Printf("Container '%s' with ID '%s'.", container.Name, container.ID)
		}
	}
}

func dockerEventListener(dockerEvents chan *docker.APIEvents, client *docker.Client) {

	err := client.AddEventListener(dockerEvents)

	if err != nil {
		logging.Error.Printf(err.Error())
		panic(err)
	}

	for event := range dockerEvents {
		if event.Status == "start" {
			if container, _ := client.InspectContainer(event.ID); container != nil {
				logging.Info.Printf("Container '%s' with ID '%s' STARTED.", container.Name, container.ID)
			}
		}
	}
}
