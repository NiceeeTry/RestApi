package service

import (
	"rest_api/pkg/repository"
	"rest_api/todo"
)

type (
	Authorization interface {
		CreateUser(user todo.User) (int, error)
		GenerateToken(username, password string) (string, error)
	}
	TodoList interface{}
	TodoItem interface{}
	Service  struct {
		Authorization
		TodoItem
		TodoList
	}
)

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
