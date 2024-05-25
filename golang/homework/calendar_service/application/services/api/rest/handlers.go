package rest

import (
	"calendar_service/application/configs"
	"calendar_service/application/services/validators"
	"calendar_service/application/usecases"
	e "calendar_service/domain/value_objects/event"
	"calendar_service/frameworks/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func Config(options *configs.Storage) {
	usecases.SetStorage(options)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event e.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error("Error in createEvent handler", zap.Error(err))
		return
	}

	err = validators.ValidateEvent(&event)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		logger.Error("Event is invalid:", zap.Error(err))
		return
	}

	err = usecases.AddEventToStorage(&event)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		logger.Error("Error in createEvent handler", zap.Error(err))
		return
	}
	msg := "Event created with ID: "
	logger.Debug(msg, zap.String("event ID: ", event.ID), zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_address", r.RemoteAddr))

	respondWithResult(w, fmt.Sprintf(msg+"%s", event.ID))
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event e.Event
	id := r.URL.Query().Get("id")
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validators.ValidateEvent(&event)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		logger.Error("Event is invalid:", zap.Error(err))
		return
	}

	err = usecases.UpdateEventFromStorage(id, &event)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := "Event updated with ID: "
	logger.Debug(msg, zap.String("event ID: ", event.ID), zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_address", r.RemoteAddr))
	respondWithResult(w, fmt.Sprintf(msg+"%s", event.ID))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := usecases.DeleteEventFromStorage(id)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := "Event deleted with ID: "
	logger.Debug(msg, zap.String("event ID: ", id), zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_address", r.RemoteAddr))
	respondWithResult(w, fmt.Sprintf(msg+"%s", id))
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get events for a specific day
	// For example, you could parse the day from the request query parameters
	day := r.URL.Query().Get("day")
	dateParams := usecases.DateParams{Day: day}
	events, err := usecases.ListEventsFromStorage("day", dateParams)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	msg := "Events for day: "
	logger.Debug(msg, events, zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_address", r.RemoteAddr))
	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write response)
	respondWithResult(w, events)
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get events for a specific week
	startOfWeek := r.URL.Query().Get("startOfWeek")
	endOfWeek := r.URL.Query().Get("endOfWeek")

	dateParams := usecases.DateParams{StartOfWeek: startOfWeek, EndOfWeek: endOfWeek}
	events, err := usecases.ListEventsFromStorage("week", dateParams)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := "Events for week: "
	logger.Debug(msg, events, zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_address", r.RemoteAddr))

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write response
	respondWithResult(w, events)
}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get events for a specific month
	month_str := r.URL.Query().Get("month")
	year_str := r.URL.Query().Get("year")
	month, err := strconv.Atoi(month_str)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	year, err := strconv.Atoi(year_str)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	dateParams := usecases.DateParams{Year: year, Month: month}
	events, err := usecases.ListEventsFromStorage("month", dateParams)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := "Events for month: "
	logger.Debug(msg, events, zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_address", r.RemoteAddr))

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Write response
	respondWithResult(w, events)
}

// TODO move to value_object
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func respondWithResult(w http.ResponseWriter, result interface{}) {
	response := Response{
		Result: result,
	}
	jsonResponse(w, response, http.StatusOK)
}

func respondWithError(w http.ResponseWriter, errMsg string, status int) {
	response := Response{
		Error: errMsg,
	}
	jsonResponse(w, response, status)
}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
