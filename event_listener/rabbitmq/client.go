package rabbitmq

import (
	"bytes"
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	uri  string
	conn *amqp.Connection
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
	}, nil
}

func (r *RabbitMQ) SerializeAndSend(channelName string, obj interface{}) error {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(obj)
	if err != nil {
		return err
	}
	r.Send(channelName, buffer.Bytes())
	return nil
}

func (r *RabbitMQ) Send(channelName string, message []byte) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		channelName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // argument
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plainapplication/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
