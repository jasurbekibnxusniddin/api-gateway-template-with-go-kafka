package dto

import "github.com/google/uuid"

type CreateTodoDto struct {
	Task string `json:"task"`
}

type UpdateTodoDto struct {
	Id          uuid.UUID `json:"id"`
	Task        string    `json:"task"`
	IsCompleted bool      `json:"is_completed"`
}
