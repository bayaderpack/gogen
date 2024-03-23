package db

import (
	// "database/sql"
	// "fmt"
	// "log"

	// "bajscheme/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

 var DB *gorm.DB

 func CreateNew(dbName string) {
	dsn := "root:@tcp(127.0.0.1:3307)/bayaderpack?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  
	// Migrate the schema
	// db.AutoMigrate(&models.Admin{})
 }
