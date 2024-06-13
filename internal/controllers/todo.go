package controllers

import (
	"bufio"
	"fmt"
	"os"

	"taskday5.2/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(ownerID uint) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var newData models.Todo

	fmt.Print("Masukkan Aktivitas: ")
	task, _ := reader.ReadString('\n')
	newData.Task = task[:len(task)-1] // Remove the newline character
	newData.Owner = ownerID

	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (tc *TodoController) UpdateTodo() (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var id uint
	var newTask string

	fmt.Print("Masukkan ID Kegiatan yang akan diupdate: ")
	fmt.Scanln(&id)
	fmt.Print("Masukkan Aktivitas Baru: ")
	newTask, _ = reader.ReadString('\n')

	todo, err := tc.model.UpdateTodo(models.Todo{ID: id, Task: newTask[:len(newTask)-1]})
	if err != nil {
		return false, err
	}
	fmt.Printf("Updated Todo: %+v\n", todo)
	return true, nil
}

func (tc *TodoController) DeleteTodo() (bool, error) {
	var id uint
	fmt.Print("Masukkan ID Kegiatan yang akan dihapus: ")
	fmt.Scanln(&id)
	err := tc.model.DeleteTodoByID(id)
	if err != nil {
		return false, err
	}
	fmt.Printf("Deleted Todo ID: %d\n", id)
	return true, nil
}

func (tc *TodoController) GetTodos(ownerID uint) ([]models.Todo, error) {
	todos, err := tc.model.GetTodosByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	fmt.Println("Todos:")
	for _, todo := range todos {
		fmt.Printf("ID: %d, Task: %s, Mark: %t\n", todo.ID, todo.Task, todo.Mark)
	}

	return todos, nil
}
