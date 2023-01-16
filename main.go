package main

// 1
import (
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/kelseyhightower/envconfig"
	// "github.com/reyhanhmdani/todolist_restAPI/configs"
	// "github.com/sirupsen/logrus"

	// kita import todocontroller nya
	"github.com/reyhanhmdani/todolist_restAPI/controllers/todoController"
	// kita import models nya
	"github.com/reyhanhmdani/todolist_restAPI/models"
)

// kita buat root
func main() {

	r := gin.Default()
	_, err := models.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	r.GET("/api.example.com/todos", todoController.Index)
	r.POST("/api.example.com/todo", todoController.Create)
	r.GET("/api.example.com/todo/:id", todoController.Show)
	r.PUT("/api.example.com/todo/:id", todoController.Update)
	r.PUT("/api.example.com/todostatus/:id", todoController.UpdateToStatus)
	r.DELETE("/api.example.com/todos", todoController.Delete)

	err = r.Run()
	if err != nil {
		log.Fatalf("failed to run route : %v", err)
	}
}
