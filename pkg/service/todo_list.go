package service

import (
	"rest_api/pkg/repository"
	"rest_api/todo"
)

type TodoListService struct {
	repo repository.TodoList
}

func (s *TodoListService) NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
