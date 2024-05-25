package dates_handlers

import (
	e "calendar_service/domain/value_objects/event"
	"fmt"
	"time"
)

func GetEventsForDay(events []*e.Event, day string) []*e.Event {
	var dayEvents []*e.Event
	targetDay, err := time.Parse("2006-01-02", day)
	if err != nil {
		fmt.Println("Error parsing day:", err)
		return dayEvents
	}
	for _, event := range events {
		startTime, err := time.Parse("2006-01-02T15:04:05Z", event.StartTime)
		if err != nil {
			fmt.Println("Error parsing start time:", err)
			continue
		}
		if startTime.Year() == targetDay.Year() &&
			startTime.Month() == targetDay.Month() &&
			startTime.Day() == targetDay.Day() {
			dayEvents = append(dayEvents, event)
		}
	}
	return dayEvents
}

func GetEventsForWeek(events []*e.Event, startOfWeek, endOfWeek string) []*e.Event {
	var weekEvents []*e.Event
	startTime, err := time.Parse("2006-01-02", startOfWeek)
	if err != nil {
		fmt.Println("Error parsing start of week:", err)
		return weekEvents
	}
	endTime, err := time.Parse("2006-01-02", endOfWeek)
	if err != nil {
		fmt.Println("Error parsing end of week:", err)
		return weekEvents
	}
	fmt.Println(events)
	for _, event := range events {
		eventStartTime, err := time.Parse("2006-01-02T15:04:05Z", event.StartTime)
		if err != nil {
			fmt.Println("Error parsing event start time:", err)
			continue
		}
		if eventStartTime.After(startTime) && eventStartTime.Before(endTime) {
			weekEvents = append(weekEvents, event)
		}
	}
	return weekEvents
}

func GetEventsForMonth(events []*e.Event, month, year int) []*e.Event {
	var monthEvents []*e.Event
	for _, event := range events {
		eventStartTime, err := time.Parse("2006-01-02T15:04:05Z", event.StartTime)
		if err != nil {
			fmt.Println("Error parsing event start time:", err)
			continue
		}
		if eventStartTime.Year() == year && int(eventStartTime.Month()) == month {
			monthEvents = append(monthEvents, event)
		}
	}
	return monthEvents
}
