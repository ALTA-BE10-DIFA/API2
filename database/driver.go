package mysql

import (
	"fmt"
	"log"

	"github.com/ALTA-BE10-DIFA/API2/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", "root", "Mihrimah220918", "localhost", 3306, "echorm")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
