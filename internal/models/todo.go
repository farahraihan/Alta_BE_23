package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Task  string `gorm:"type:varchar(100)"`
	Mark  bool
	Owner uint // Assuming Owner is a field to relate to user
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(db *gorm.DB) *TodoModel {
	return &TodoModel{db: db}
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.db.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}
	return newData, nil
}

func (tm *TodoModel) UpdateTodo(updatedData Todo) (Todo, error) {
	err := tm.db.Save(&updatedData).Error
	if err != nil {
		return Todo{}, err
	}
	return updatedData, nil
}

func (tm *TodoModel) DeleteTodoByID(id uint) error {
	err := tm.db.Delete(&Todo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (tm *TodoModel) GetTodos() ([]Todo, error) {
	var todos []Todo
	err := tm.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tm *TodoModel) GetTodosByOwner(ownerID uint) ([]Todo, error) {
	var todos []Todo
	err := tm.db.Where("owner = ?", ownerID).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
