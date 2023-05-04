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
		GetAll(userId int) ([]todo.TodoList, error)
		GetById(userId, listId int) (todo.TodoList, error)
		Delete(userId int, listId int) error
		Update(userId, id int, input todo.UpdateListInput) error
	}
	TodoItem interface {
		Create(userId, listId int, input todo.TodoItem) (int, error)
		GetAll(userId, listId int) ([]todo.TodoItem, error)
	}
	Service struct {
		Authorization
		TodoItem
		TodoList
	}
)

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
