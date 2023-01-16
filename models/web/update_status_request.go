package web

type UpdateStatusTodo struct {
	//gorm.Model
	Id     int64 `json:"id"`
	Status bool  `json:"status" binding:"required"`
}
