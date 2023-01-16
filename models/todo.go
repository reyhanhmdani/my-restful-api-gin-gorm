package models

type Todo struct {
	//gorm.Model
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"type:varchar(300)" json:"title" binding:"required,min=2"`
	Status bool   `gorm:"default:false" json:"status"`
}
