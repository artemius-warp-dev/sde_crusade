package queue

type EventQueue interface {
	Publish(queueName string, body []byte) error
	Consume(queueName string) (<-chan []byte, error)
	Close() error
	GetQueue() string
}
