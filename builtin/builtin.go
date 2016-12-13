package builtin

import (
	_ "github.com/dmportella/docker-beat/builtin/consumers/rabbitmq" // only imported to run the package init
	_ "github.com/dmportella/docker-beat/builtin/consumers/webhook"  // only imported to run the package init
)
