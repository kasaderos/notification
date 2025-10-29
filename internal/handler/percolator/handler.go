package percolator

type Handler struct {
	service  EventService
	consumer Consumer
}

func New(service EventService, consumer Consumer) *Handler {
	return &Handler{service: service, consumer: consumer}
}
