package db

import (
	// "database/sql"
	// "fmt"
	// "log"

	// "bajscheme/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

 var DB *gorm.DB
 var err error

 func init() {
	dsn := "root:@tcp(127.0.0.1:3307)/bayaderpack?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	
 }
 func CreateNew(dbName string) {
	dsn := "root:@tcp(127.0.0.1:3307)/bayaderpack?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
  
	// Migrate the schema
	// db.AutoMigrate(&models.Admin{})
 }
