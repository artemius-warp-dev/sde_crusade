package queues

import (
	"calendar_service/application/configs"
	"calendar_service/application/services/jobs"
	"calendar_service/application/usecases"
	"calendar_service/frameworks/logger"
	queue_impl "calendar_service/frameworks/queue"
)

func Init(config *configs.ServerConfig) {
	var err error
	//TODO add type of queue. For mock without rabbitmq
	usecases.SetStorage(&config.Storage)
	usecases.EventQueue, err = queue_impl.NewRabbitMQBroker(*config.Queue.ConnStr, config.Queue.QueueName)
	if err != nil {
		logger.Error("Failed to initilize message broker: ", err)
	}

	logger.Info("Message broker is ready")
}

func Schedule() {
	go jobs.Scheduler()
	logger.Info("Scheduler running")
}

func Consume() {
	go jobs.Sender()
	logger.Info("Consumer running")
}
