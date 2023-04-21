package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
