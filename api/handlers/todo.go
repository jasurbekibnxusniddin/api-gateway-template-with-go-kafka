package handlers

import (
	"github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/dto"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateTodo(c echo.Context) error {

	reqBody := dto.CreateTodoDto{}

	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	if err := h.event.CreateTodo(&reqBody); err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(201, reqBody)
}
