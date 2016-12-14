package plugin

// EventConsumer is the interface for all the consumers used by docker-beat.
type EventConsumer interface {
	OnEvent(events DockerEvent)
}

var (
	consumers = []EventConsumer{}
)

// RegisterConsumer adds the consumer to the registry.
func RegisterConsumer(consumer EventConsumer) {
	consumers = append(consumers, consumer)
}

// GetConsumer temp method for testing
func GetConsumer() (consumer EventConsumer) {
	return consumers[0]
}
