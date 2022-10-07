package database

import (
	"assignment2/item"
	"assignment2/order"
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=assignment2 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")

	db.Migrator().DropTable(
		&item.Item{},
		&order.Order{},
	)

	db.AutoMigrate(
		&item.Item{},
		&order.Order{},
	)

	fmt.Println("successfully migrating table")
}

func GetDB() *gorm.DB {
	return db
}
