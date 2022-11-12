package database

import (
	"FP2/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBinit(dbHost, dbPort, dbUsername, dbPassword, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbName)
	//root:iNdkC2NQNZ2bdQJhmwnE@containers-us-west-119.railway.app:7261/railway
	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR: Failed to connect to Database -> %v\n", err)
	}
	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return db
}
