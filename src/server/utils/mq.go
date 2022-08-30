package utils

import (
	"context"
	"fmt"
	"opengin/server/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	ExchangeType_Direct  = "direct"
	ExchangeType_Fanout  = "fanout"
	ExchangeType_Topic   = "topic"
	ExchangeType_Headers = "headers"
)

type RabbitMqOptions struct {
	ExchangeName string
	QueueName    string
	RoutingKey   string
}

type RabbitMqHandler struct {
	MqOptions  *RabbitMqOptions
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue
}

type MessageHandler func(msg []byte)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}

func NewRabbitMqHandler(options *RabbitMqOptions) *RabbitMqHandler {
	host := config.Settings.RabbitMQ.Host
	port := config.Settings.RabbitMQ.Port
	user := config.Settings.RabbitMQ.User
	password := config.Settings.RabbitMQ.Password
	virtualHost := config.Settings.RabbitMQ.VirtualHost

	conn, err := amqp.Dial("amqp://" + user + ":" + password + "@" + host + ":" + port + "/" + virtualHost)
	failOnError(err, "Failed to establish a connection.")

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel.")

	args := amqp.Table{
		"x-message-ttl": 1000 * 60 * 60 * 24,
	}

	err = channel.ExchangeDeclare(options.ExchangeName, ExchangeType_Fanout, true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange.")

	queue, err := channel.QueueDeclare(options.QueueName, true, false, false, false, args)
	failOnError(err, "Failed to declare a queue.")

	channel.QueueBind(queue.Name, options.RoutingKey, options.ExchangeName, false, nil)

	return &RabbitMqHandler{
		MqOptions:  options,
		Connection: conn,
		Channel:    channel,
		Queue:      &queue,
	}
}

func (r *RabbitMqHandler) Publish(message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.Channel.PublishWithContext(ctx, r.MqOptions.ExchangeName, r.MqOptions.RoutingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(message),
	})
	failOnError(err, "Failed to publish a message.")
}

func (r *RabbitMqHandler) StartConsuming(handler MessageHandler) {
	r.Channel.Qos(1, 0, false)
	deliveries, err := r.Channel.Consume(r.MqOptions.QueueName, "", false, false, false, false, nil)
	failOnError(err, "Failed to consume messages.")
	var forever chan struct{}

	go func() {
		for delivery := range deliveries {
			handler(delivery.Body)
			delivery.Ack(false)
		}
	}()

	fmt.Println("Consuming messages...")
	<-forever
}

func (r *RabbitMqHandler) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}

	if r.Connection != nil {
		r.Connection.Close()
	}
}
