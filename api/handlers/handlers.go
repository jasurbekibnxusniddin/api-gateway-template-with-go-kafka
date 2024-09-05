package handlers

import "github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/event"

type handler struct {
	event event.EventI
}

func NewHandler(e event.EventI) *handler {
	return &handler{
		event: e,
	}
}
