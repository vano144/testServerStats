package statistics

import "fmt"

type EntityDoesNotAlreadyExists struct {
	modelName string
	id int
}

func (e *EntityDoesNotAlreadyExists) Error() string {
	return fmt.Sprintf("entity %s doesn't exists with id: %d", e.modelName, e.id)
}

type EntityAlreadyExists struct {
	modelName string
	id int
}

func (e *EntityAlreadyExists) Error() string {
	return fmt.Sprintf("entity %s already exists with id: %d", e.modelName, e.id)
}

