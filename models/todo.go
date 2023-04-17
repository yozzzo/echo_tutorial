package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `json:"content"`
}

func CreateTodo(content string) (Todo, error) {
	var todo Todo = Todo{Content: content}
	db.Create(&todo)
	return todo, err
}

func GetTodo(id int) (Todo, error) {
	var todo Todo
	db.First(&todo, id)
	return todo, err
}

func GetAllTodo() ([]Todo, error) {
	var todos []Todo
	db.Find(&todos)
	return todos, err
}

func UpdateTodo(todo Todo) (Todo, error) {
	db.Save(&todo)
	return todo, err
}

func DeleteTodo(todo Todo) error {
	db.Delete(&todo)
	return err
}

// メモ
// Named return value使わない。
// 更新した時に、フロントで何に更新されたかわかるように、更新したデータを返す。
