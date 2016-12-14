package main

import (
	"flag"
	"fmt"
	_ "github.com/dmportella/docker-beat/builtin"
	"github.com/dmportella/docker-beat/logging"
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
	Consumer       string
	Version        bool
	Verbose        bool
	Help           bool
)

func init() {
	const (
		defaultDockerEndpoint = "unix:///var/run/docker.sock"
		dockerEndpointUsage   = "The Url or unix socket address for the Docker Remote API."
		defaultConsumer       = "console"
		ConsumerUsage         = "Consumer to use: Webhook, Rabbitmq, etc."
	)

	flag.StringVar(&DockerEndpoint, "docker-endpoint", defaultDockerEndpoint, dockerEndpointUsage)
	flag.StringVar(&Consumer, "consumer", defaultConsumer, ConsumerUsage)

	const (
		defaultHelp    = false
		helpUsage      = "Prints the help information."
		defaultVerbose = false
		verboseUsage   = "Redirect trace information to the standard out."
	)

	flag.BoolVar(&Verbose, "verbose", defaultVerbose, verboseUsage)
	flag.BoolVar(&Help, "help", defaultHelp, helpUsage)
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

	if Help {
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

	beat, err := newDockerBeat(DockerEndpoint, Consumer)

	if err == nil {
		beat.Start()
	} else {
		panic(err)
	}

	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}
