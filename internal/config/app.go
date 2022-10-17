package config

import (
	"fmt"

	"github.com/md/go-pro-main/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var urlDSN = "root:22031998@tcp(localhost:3306)/go_test?parseTime=true"

func Connect() {

}

func DataMigration() {
	DB, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	// mysql.Open(urlDSN), &gorm.Config{}

	if err != nil {
		fmt.Print(err.Error())
		panic("connetion failed...")
	}

	DB.AutoMigrate(&models.Post{})
}
