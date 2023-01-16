package web

type UpdateTodo struct {
	//gorm.Model
	Id     int64  `json:"id"`
	Title  string `json:"title" binding:"required,min=2"`
	Status bool   `json:"status"`
}
