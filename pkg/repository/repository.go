package repository

import (
	"rest_api/todo"

	"github.com/jmoiron/sqlx"
)

type (
	Authorization interface {
		CreateUser(user todo.User) (int, error)
		GetUser(username, password string) (todo.User, error)
	}
	TodoList interface {
		Create(userId int, list todo.TodoList) (int, error)
		GetAll(userId int) ([]todo.TodoList, error)
		GetById(userId, listId int) (todo.TodoList, error)
		Delete(userId int, listId int) error
		Update(userId, id int, input todo.UpdateListInput) error
	}
	TodoItem interface {
		Create(listId int, item todo.TodoItem) (int, error)
		GetAll(userId, listId int) ([]todo.TodoItem, error)
	}
	Repository struct {
		Authorization
		TodoItem
		TodoList
	}
)

func NewRepository(DB *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(DB),
		TodoList:      NewTodoListPostgre(DB),
		TodoItem:      NewTodoItemPostgres(DB),
	}
}
