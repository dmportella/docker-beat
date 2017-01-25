package plugin

// EventConsumer is the interface for all the consumers used by docker-beat.
type EventConsumer interface {
	OnEvent(events DockerEvent)
}

var (
	consumers map[string]EventConsumer
)

func init() {
	consumers = make(map[string]EventConsumer)
}

// RegisterConsumer Adds a consumer to the registry.
func RegisterConsumer(name string, consumer EventConsumer) {
	consumers[name] = consumer
}

// GetConsumer Returns the selected consumer
func GetConsumer(name string) (consumer EventConsumer) {
	return consumers[name]
}
