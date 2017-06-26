package main

import "fmt"

var currentID int
var todos Todos

func init() {
	CreateTodo(Todo{Name: "Write Presentation"})
	CreateTodo(Todo{Name: "Host Meetup"})
}

func FindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}

	//if no todo found for given ID: return empty
	return Todo{}
}

func CreateTodo(t Todo) Todo {
	currentID += 1
	t.Id = currentID
	todos = append(todos, t)
	return t
}

func DestroyTodod(id int) error {
	for i, t := range todos {
		if id == t.Id {
			todos = append(todos[:id], todos[i+1:]...)
			return nil
		}
	}

	//if no todo found for given ID: throw error
	return fmt.Errorf("Could not find Todo with id %d to delete", id)
}