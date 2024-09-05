package api

import (
	"github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/api/handlers"
	"github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/event"
	"github.com/labstack/echo/v4"
)

type Options struct {
	Event event.EventI
}

func Api(o Options) *echo.Echo {

	h := handlers.NewHandler(o.Event)

	router := echo.New()

	api := router.Group("/api")
	{
		api.POST("/todo", h.CreateTodo)
	}

	return router
}
