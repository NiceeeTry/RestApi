package service

import "rest_api/pkg/repository"

type (
	Authorization interface{}
	TodoList      interface{}
	TodoItem      interface{}
	Service       struct {
		Authorization
		TodoItem
		TodoList
	}
)

func NewService(repos repository.Repository) *Service {
	return &Service{}
}
