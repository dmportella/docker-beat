package plugins

// EventConsumer is the interface for all the consumers used by docker-beat.
type EventConsumer interface {
	OnEvent(events DockerEvent)
}
