package web

type TodoCreate struct {
	//gorm.Model
	Title  string `json:"title" binding:"required,min=2"`
	Status bool   `json:"status" binding:"required"`
}
