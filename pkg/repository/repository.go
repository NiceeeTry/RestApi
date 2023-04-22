package repository

import "github.com/jmoiron/sqlx"

type (
	Authorization interface{}
	TodoList      interface{}
	TodoItem      interface{}
	Repository    struct {
		Authorization
		TodoItem
		TodoList
	}
)

func NewRepository(DB *sqlx.DB) *Repository {
	return &Repository{}
}
