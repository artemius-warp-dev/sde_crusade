package queue

import (
	"calendar_service/frameworks/logger"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queueName string
}

func NewRabbitMQBroker(url string, queueName string) (*RabbitMQBroker, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Error("Failed to connect to RabitMQ: ", err)
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		logger.Error("Failed to open a channel: ", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when used
		false, // exclusive
		false, // no-wait
		nil,   //arguments
	)

	return &RabbitMQBroker{conn: conn, ch: ch, queueName: queueName}, nil
}

func (r *RabbitMQBroker) Publish(queueName string, body []byte) error {

	err := r.ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     //mandatory
		false,     //immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err

}

func (r *RabbitMQBroker) Consume(queueName string) (<-chan []byte, error) {
	msgs, err := r.ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)
	go func() {
		for d := range msgs {
			out <- d.Body
		}
		close(out)

	}()

	return out, nil
}

func (r *RabbitMQBroker) Close() error {
	if err := r.ch.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}

func (r *RabbitMQBroker) GetQueue() string {
	return r.queueName
}
