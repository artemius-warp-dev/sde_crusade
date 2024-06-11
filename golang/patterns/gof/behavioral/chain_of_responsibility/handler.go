package main

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(service DateTemperatureService)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {

	h.next = handler
	return handler

}

func (h *BaseHandler) Handle(service DateTemperatureService) {

	if h.next != nil {
		h.next.Handle(service)
	}
}
