package serializer

import model "project2019/models"

type Todo struct {
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	Info string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}


// BuildTodo 序列化待办事项
func BuildTodo(item model.TodoList) Todo {
	return Todo{
		Title:    item.Title,
		Status:    item.Status,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}