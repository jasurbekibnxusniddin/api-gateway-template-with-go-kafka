package event

import (
	"encoding/json"
	"log"

	"github.com/jasurbekibnxusniddin/api-gateway-template-with-go-kafka/dto"
)

type EventI interface {
	CreateTodo(data *dto.CreateTodoDto) error
}
type event struct {
}

func NewEvent() EventI {
	return &event{}
}

func (e *event) CreateTodo(data *dto.CreateTodoDto) error {

	topic := "create_todo"

	msg, err := json.Marshal(data)
	if err != nil {
		log.Println("error on Marshal. Error:", err)
		return err
	}

	err = SendMessage(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

func (e *event) UpdateTodo(data *dto.UpdateTodoDto) error {

	topic := "update_todo"

	msg, err := json.Marshal(data)
	if err != nil {
		log.Println("error on Marshal. Error:", err)
		return err
	}

	err = SendMessage(topic, msg)
	if err != nil {
		return err
	}

	return nil
}
