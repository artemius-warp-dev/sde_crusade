package webservers

import (
	"calendar_service/application/configs"
	"calendar_service/application/services/api/rest"
	"calendar_service/frameworks/logger"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Server(c *configs.ServerConfig) {

	rest.Config(&c.Storage)
	if err := http.ListenAndServe(c.HTTPAddress, nil); err != nil {
		fmt.Println(err)
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func RouterInit(c configs.ServerConfig) {

	http.HandleFunc("/create_event", applyMiddleware(rest.CreateEvent, logMiddleware))
	http.HandleFunc("/update_event", applyMiddleware(rest.UpdateEvent, logMiddleware))
	http.HandleFunc("/delete_event", applyMiddleware(rest.DeleteEvent, logMiddleware))
	http.HandleFunc("/events_for_day", applyMiddleware(rest.EventsForDay, logMiddleware))
	http.HandleFunc("/events_for_week", applyMiddleware(rest.EventsForWeek, logMiddleware))
	http.HandleFunc("/events_for_month", applyMiddleware(rest.EventsForMonth, logMiddleware))

}

func applyMiddleware(next http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		next = m(next)
	}
	return next
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next(w, r)

		logger.Info(fmt.Sprintf("%s/%s/%s/%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start)))
	}

}
