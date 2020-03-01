package models

import "github.com/jinzhu/gorm"

// TodoList 模型
type TodoList struct {
	gorm.Model
	Title  string
	Status bool
	Info   string
	User User `gorm:"ForeignKey:Username"`
	Refer string
}
