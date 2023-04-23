package service

import (
	"rest_api/pkg/repository"
	"rest_api/todo"
)

type (
	Authorization interface {
		CreateUser(user todo.User) (int, error)
		GenerateToken(username, password string) (string, error)
		ParseToken(token string) (int, error)
	}
	TodoList interface {
		Create(userId int, list todo.TodoList) (int, error)
	}
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
