package main

import (
	"github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/api"
	"github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/event"
)

func main() {

	event := event.NewEvent()

	router := api.Api(
		api.Options{
			Event: event,
		},
	)

	router.Start(":8080")
}
