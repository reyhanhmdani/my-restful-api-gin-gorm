package todoController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/reyhanhmdani/todolist_restAPI/helper"
	"github.com/reyhanhmdani/todolist_restAPI/models"
	"github.com/reyhanhmdani/todolist_restAPI/models/web"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Index(prm *gin.Context) {
	var todos []models.Todo
	// kemudian kita mengambil semua data di todos yang ada di database terus kita masukkan variablenya dengan pointer
	db, err := models.ConnectDB()
	if err != nil {
		prm.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Fatal("error")
	}
	if err := db.Find(&todos).Error; err != nil {
		logrus.Error("ga masuk ke database")
		prm.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	// kita akan mengirim todos ini ke response JSON
	logrus.Info("Semua data muncul")
	// 23
	prm.JSON(http.StatusOK, gin.H{"todos": todos})
}

func Show(prm *gin.Context) {
	//
	var todos models.Todo

	//
	// kita ambil id nya dulu
	id := prm.Param("id")

	//
	// kita query
	db, err := models.ConnectDB()
	helper.Error(err)

	if err := db.First(&todos, id).Error; err != nil {
		// apa bila eror atau sama dengan nil kita switch error nya , kita cek error nya apakah error nya datanya tidak
		// di temukan atau karna internal server error
		switch err {
		case gorm.ErrRecordNotFound:
			// kita akan mengembalikan response data tidak di temukan
			logrus.Info("Data nya ga ada")
			// ini apa bila kalau kita mengambil data yang tidak ada maka akan muncul di bawah ini
			prm.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ada"})
			// kita return supaya prosesnya terhenti setelah kita mengirim JSON  di atas ini
			return
			// 30
		default:
			prm.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	logrus.Info(todos)

	// jika tidak error saya ambil respon dari todos yang saya ambil dari database
	prm.JSON(http.StatusOK, gin.H{"todo": todos})
}

func Create(prm *gin.Context) {

	var input web.TodoCreate

	// var todos models.Todo

	if err := prm.ShouldBindJSON(&input); err != nil {
		prm.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logrus.Error(err.Error())
		// kita return supaya prosesya terhenti setelah kita mengirim json nya
		return
		// jika jsonbody nya berhasil kita ambil atau kita simpan ke struct T0do maka di bawah sini datanya kita simpan ke database
	}

	db, err := models.ConnectDB()
	if err != nil {
		helper.Error(err)
	}

	// create todo
	todo := models.Todo{Title: input.Title, Status: input.Status}

	if err := db.Create(&todo).Error; err != nil {
		logrus.Error("datanya tidak ada")
		prm.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// // kalau berhasil kita kasih respon
	prm.JSON(http.StatusOK, gin.H{"todos": todo})

	prm.JSON(http.StatusCreated, gin.H{"message": "data berhasil terbuat"})

	//prm.JSON(http.StatusOK, gin.H{"message": &todos})

	logrus.Info("data sudah terbuat")

}

func Update(prm *gin.Context) {
	var todos models.Todo

	// kita ambil id nya dulu
	//id := prm.Param("id")

	if err := prm.ShouldBindJSON(&todos); err != nil {
		prm.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logrus.Error(err.Error())
		// kita return supaya prosesya terhenti setelah kita mengirim json nya
		return
		// jika jsonbody nya berhasil kita ambil atau kita simpan ke struct T0do maka di bawah sini datanya kita simpan ke database // maka kita update ke database
	}

	db, err := models.ConnectDB()
	helper.Error(err)

	if db.Model(&todos).Where("id=?", prm.Param("id")).First(&todos).RowsAffected == 0 {
		//logrus.Error("data gagal terupdate")
		prm.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "tidak dapat mengupdate todo"})
		// kita return supaya prosesya terhenti setelah kita mengirim json nya
		return
	}

	// jika t0do nya berhasil di update kita kembalikan response json
	prm.JSON(http.StatusOK, gin.H{"message": "Data berhasil di Update"})
	logrus.Info("data berhasil terupdate")

}

func UpdateToStatus(prm *gin.Context) {
	//var todos web.UpdateStatusTodo

	var todos models.Todo

	db, err := models.ConnectDB()
	if err != nil {
		helper.Error(err)
	}

	if err := db.Where("id = ?", prm.Param("id")).First(&todos).Error; err != nil {
		logrus.Error("ID NOT FOUND!")
		prm.JSON(http.StatusBadRequest, gin.H{"error": "ID not found!"})
		return
	}

	// Validate input
	var input web.UpdateStatusTodo
	if err := prm.ShouldBindJSON(&input); err != nil {
		prm.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{Status: input.Status}

	db.Where("id", prm.Param("id")).Updates(&todo)
	// jika t0do nya berhasil di update kita kembalikan response json
	prm.JSON(http.StatusOK, gin.H{"message": "Data berhasil di Update"})
	logrus.Info("data berhasil terupdate")

}

func Delete(prm *gin.Context) {
	var todos models.Todo

	var input struct {
		Id json.Number
	}
	if err := prm.ShouldBindJSON(&input); err != nil {
		prm.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logrus.Error(err.Error())
		// kita return supaya prosesya terhenti setelah kita mengirim json nya
		return
	}

	// kita akan mengambil item id value nya dalam bentulk int 64
	id, _ := input.Id.Int64()
	// proses delete data
	db, err := models.ConnectDB()
	helper.Error(err)

	if db.Delete(&todos, id).RowsAffected == 0 {
		prm.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "tidak dapat menghapus"})
		return
	}
	// jika berhasil maka kita kembalikan response
	logrus.Info(http.StatusOK)
	prm.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus Data"})
	logrus.Info("Data sudah terhapus")
}
