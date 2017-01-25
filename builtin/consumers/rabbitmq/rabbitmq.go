package rabbitmq

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dmportella/docker-beat/logging"
	"github.com/dmportella/docker-beat/plugin"
	"github.com/streadway/amqp"
)

var (
	rabbitmqEndpoint     string
	rabbitmqExchange     string
	rabbitmqExchangeType string
	rabbitmqRoutingKey   string
	rabbitmqReliable     bool
)

const (
	defaultRabbitmqEndpoint = "amqp://guest:guest@localhost:5672/"
	rabbitmqEndpointUsage   = "rabbitmq: The URL that events will be published too."

	defaultRabbitmqExchange = "docker-beat"
	rabbitmqExchangeUsage   = "rabbitmq: The exchange docker-beat will publish messages too."

	defaultRabbitmqExchangeType = "fanout"
	rabbitmqExchangeTypeUsage   = "rabbitmq: The exchange type that docker-beat will create/connect too. (direct|fanout|topic|x-custom)"

	defaultRabbitmqRoutingKey = "docker-event"
	rabbitmqRoutingKeyUsage   = "rabbitmq: The routing key for messages published to the exchange. (default: docker-event"

	defaultRabbitmqReliable = false
	rabbitmqReliableUsage   = "rabbitmq: The ensures messages published are confirmed."
)

type consumer struct {
	Debug        bool
	Indent       bool
	endpoint     string
	exchange     string
	exchangeType string
	routingKey   string
	reliable     bool
	connection   *amqp.Connection
}

func (consumer *consumer) OnEvent(event plugin.DockerEvent) {
	var data []byte

	if consumer.Indent {
		data, _ = json.MarshalIndent(event, "", "    ")
	} else {
		data, _ = json.Marshal(event)
	}

	err := consumer.publish(data)
	if err != nil {
		logging.Error.Printf("Error publishing event '%s'\n", err)
	}
}

func init() {
	flag.StringVar(&rabbitmqEndpoint, "rabbitmq-endpoint", defaultRabbitmqEndpoint, rabbitmqEndpointUsage)
	flag.StringVar(&rabbitmqExchange, "rabbitmq-exchange", defaultRabbitmqExchange, rabbitmqExchangeUsage)
	flag.StringVar(&rabbitmqExchangeType, "rabbitmq-exchange-type", defaultRabbitmqExchangeType, rabbitmqExchangeTypeUsage)
	flag.StringVar(&rabbitmqRoutingKey, "rabbitmq-routing-key", defaultRabbitmqRoutingKey, rabbitmqRoutingKeyUsage)
	flag.BoolVar(&rabbitmqReliable, "rabbitmq-reliable", defaultRabbitmqReliable, rabbitmqReliableUsage)

	consumer := &consumer{
		endpoint:     rabbitmqEndpoint,
		exchange:     rabbitmqExchange,
		exchangeType: rabbitmqExchangeType,
		routingKey:   rabbitmqRoutingKey,
		reliable:     rabbitmqReliable,
	}

	plugin.RegisterConsumer("rabbitmq", consumer)
}

func (consumer *consumer) publish(body []byte) error {
	if consumer.connection == nil {
		consumer.endpoint = rabbitmqEndpoint
		consumer.exchange = rabbitmqExchange
		consumer.exchangeType = rabbitmqExchangeType
		consumer.routingKey = rabbitmqRoutingKey
		consumer.reliable = rabbitmqReliable

		connection, err := amqp.Dial(consumer.endpoint)
		if err != nil {
			return fmt.Errorf("Dial: %s", err)
		}

		consumer.connection = connection
	}

	channel, err := consumer.connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	if err := channel.ExchangeDeclare(
		consumer.exchange,     // name
		consumer.exchangeType, // type
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // noWait
		nil,   // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if consumer.reliable {
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("Channel could not be put into confirm mode: %s", err)
		}

		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		defer confirmOne(confirms)
	}

	if err = channel.Publish(
		consumer.exchange,   // publish to an exchange
		consumer.routingKey, // routing to 0 or more queues
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	logging.Info.Printf("Waiting for confirmation of one publishing.")

	if confirmed := <-confirms; confirmed.Ack {
		logging.Info.Printf("Confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		logging.Error.Printf("Failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
