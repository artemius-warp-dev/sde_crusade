package jobs

import (
	"calendar_service/application/usecases"
	e "calendar_service/domain/value_objects/event"
	"calendar_service/frameworks/logger"
	"encoding/json"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func Sender(counter prometheus.Counter) {

	msgs, err := usecases.ConsumeEvents()

	if err != nil {
		logger.Error("Failed to consume messages: ", err)

	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var event e.Event
			err := json.Unmarshal(d, &event)
			if err != nil {
				logger.Error("Failed to unmarshal message: ", err)
				continue
			}
			fmt.Println("Received event: ", event)
			err = usecases.MarkEventAsNotified(event.ID)
			if err != nil {
				logger.Error("Failed to mark events as notified: ", err)
			}
			counter.Inc()
		}
	}()

	logger.Info("Waiting for messages in queue: ", usecases.EventQueue.GetQueue())
	<-forever
}
