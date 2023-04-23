package repository

import (
	"rest_api/todo"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgre struct {
	db *sqlx.DB
}

func NewTodoListPostgre(db *sqlx.DB) *TodoListPostgre {
	return &TodoListPostgre{db: db}
}

func (r *TodoListPostgre) Create(userId int, list todo.TodoList) (int, error) {
}
