package jobs

import (
	"calendar_service/frameworks/logger"

	"calendar_service/application/usecases"

	"github.com/jasonlvhit/gocron"
)

func Scheduler() {
	gocron.Every(30).Second().Do(task)
	<-gocron.Start()
}

func task() {
	logger.Info("Scheduler task start")
	events, err := usecases.FetchEventsForNotification()
	if err != nil {
		logger.Error("Error fetching events: ", err)
		return
	}

	if len(events) == 0 {
		logger.Info("No new events to notify")
		return
	} else {
		logger.Info("Events has been fetched")
	}

	for _, event := range events {
		err = usecases.PublishEvent(event)
		if err != nil {
			logger.Error("Failed to publish event: ", err)
		}
	}

}
