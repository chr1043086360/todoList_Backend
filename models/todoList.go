package models

import "github.com/jinzhu/gorm"

// TodoList 模型
type TodoList struct {
	gorm.Model
	Id     uint
	Title  string
	Status bool
	Info   string
}
