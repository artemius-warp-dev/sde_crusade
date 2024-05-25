package grpc

import (
	"calendar_service/application/configs"
	"calendar_service/application/services/validators"
	"calendar_service/application/usecases"
	"calendar_service/domain/value_objects/event"
	"calendar_service/frameworks/logger"
	pb "calendar_service/frameworks/protobuf/events"
	"context"
	"fmt"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type eventServer struct {
	pb.UnimplementedEventServiceServer
}

func Config(options *configs.Storage) {
	usecases.SetStorage(options)
}

// Implement the gRPC methods

func InitEventServer(srv *grpc.Server) {
	pb.RegisterEventServiceServer(srv, &eventServer{})
}

func (s *eventServer) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	// Implement logic to create an event
	logger.Info("Received CreateEvent request:", req)
	event := fetchEventFromReq(req.Event)
	err := validators.ValidateEvent(event)
	if err != nil {
		logger.Error("Event in grpc is invalid:", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = usecases.AddEventToStorage(event)
	if err != nil {
		logger.Error("Error in  grpc createEvent handler", zap.Error(err))
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.CreateEventResponse{
		Success: true,
		Message: fmt.Sprintf("Event %s created successfully", event.ID),
	}, nil
}

func (s *eventServer) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	// Implement logic to update an event
	logger.Info("Received UpdateEvent request:", req)
	id := req.Id
	event := fetchEventFromReq(req.Event)

	err := validators.ValidateEvent(event)
	if err != nil {
		logger.Error("Event in grpc is invalid:", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = usecases.UpdateEventFromStorage(id, event)
	if err != nil {
		logger.Error("Error in  grpc UpdateEvent handler", zap.Error(err))
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.UpdateEventResponse{
		Success: true,
		Message: "Event updated successfully",
	}, nil
}

func (s *eventServer) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	// Implement logic to delete an event
	logger.Info("Received DeleteEvent request:", req)

	id := req.Id
	err := usecases.DeleteEventFromStorage(id)
	if err != nil {
		logger.Error("Error in  grpc DeleteEvent handler", zap.Error(err))
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.DeleteEventResponse{
		Success: true,
		Message: "Event deleted successfully",
	}, nil
}

func (s *eventServer) EventsForDay(ctx context.Context, req *pb.EventsForDayRequest) (*pb.EventsForDayResponse, error) {
	// Implement logic to fetch events for a day
	logger.Info("Received EventsForDay request:", req)

	day := req.GetDay()
	dateParams := usecases.DateParams{Day: day}
	events, err := usecases.ListEventsFromStorage("day", dateParams)
	if err != nil {
		logger.Error("Error in  grpc EventsForDay handler", zap.Error(err))
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.EventsForDayResponse{
		Events: convertToPBEvents(events), // Implement logic to fetch events
	}, nil
}

func (s *eventServer) EventsForWeek(ctx context.Context, req *pb.EventsForWeekRequest) (*pb.EventsForWeekResponse, error) {
	// Implement logic to fetch events for a day
	logger.Info("Received EventsForWeek request:", req)
	startOfWeek := req.GetStartOfWeek()
	endOfWeek := req.GetEndOfWeek()
	dateParams := usecases.DateParams{StartOfWeek: startOfWeek, EndOfWeek: endOfWeek}
	events, err := usecases.ListEventsFromStorage("week", dateParams)
	if err != nil {
		logger.Error("Error in  grpc EventsForWeek handler", zap.Error(err))
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.EventsForWeekResponse{
		Events: convertToPBEvents(events), // Implement logic to fetch events
	}, nil
}

func (s *eventServer) EventsForMonth(ctx context.Context, req *pb.EventsForMonthRequest) (*pb.EventsForMonthResponse, error) {
	// Implement logic to fetch events for a day
	logger.Info("Received EventsForMonth request:", req)
	month, err := strconv.Atoi(req.GetMonth())
	year, err := strconv.Atoi(req.GetYear())
	dateParams := usecases.DateParams{Month: month, Year: year}
	events, err := usecases.ListEventsFromStorage("day", dateParams)
	if err != nil {
		logger.Error("Error in  grpc EventsForMonth handler", zap.Error(err))
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.EventsForMonthResponse{
		Events: convertToPBEvents(events), // Implement logic to fetch events
	}, nil
}

func fetchEventFromReq(e *pb.Event) *event.Event {
	// Extract data from the request message
	event := &event.Event{
		ID:          e.GetId(),
		Title:       e.GetTitle(),
		Description: e.GetDescription(),
		StartTime:   e.GetStartTime(),
		EndTime:     e.GetEndTime(),
		Location:    e.GetLocation(),
		Attendees:   e.GetAttendees(),
	}
	return event
}

func convertToPBEvents(events []*event.Event) []*pb.Event {
	pbEvents := make([]*pb.Event, len(events))

	for i, event := range events {
		pbEvents[i] = &pb.Event{
			Id:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			Location:    event.Location,
			Attendees:   event.Attendees,
			// Convert other fields as needed
		}
	}

	return pbEvents
}
