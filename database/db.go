package database

import (
	"FP2/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBinit(dbHost, dbPort, dbUsername, dbPassword, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	//root:qF0Z1vu4M7yoJHdil1ch@containers-us-west-120.railway.app:7457/railway
	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR: Failed to connect to Database -> %v\n", err)
	}
	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return db
}
