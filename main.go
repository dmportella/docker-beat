package main

import (
	"flag"
	"fmt"
	"github.com/dmportella/docker-beat/logging"
	"io/ioutil"
	"os"
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
	Version bool
	Verbose bool
)

func init() {
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

	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}
