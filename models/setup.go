package models

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/reyhanhmdani/todolist_restAPI/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectDB() (*gorm.DB, error) {

	var cfg configs.Config

	// pr
	err1 := envconfig.Process("", &cfg)
	if err1 != nil {
		logrus.Fatal("error")
	}
	fmt.Printf("%+v\n ", cfg)
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	Database, err := gorm.Open(mysql.Open(dsn))
	// jika error selama proses penyambungan ke koneksi maka akan error
	if err != nil {
		return nil, err
	}
	err = Database.AutoMigrate(&Todo{})
	if err != nil {
		return nil, err
	}

	return Database, err

}
