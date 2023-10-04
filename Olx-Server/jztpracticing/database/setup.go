package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect_to(connect string) {

	// the gorm.open used to create a database connection ,it takes two arguments
	//the first arg is to specify to which database driver should be connected
	//the second arg is to manage gorm's behaviour in this case its emty so it will use its default cofiguration

	DB, err = gorm.Open(postgres.Open(connect), &gorm.Config{})

	if err != nil {
		panic("cannot connect to the databse...")
	}

	fmt.Println("connected to the databse successfully ... ")
}

// variadic function to receive all the structs to migrate to a table which implemented the Group_tables interface

func Migrte_all(models ...Group_tables) {
	for _, model := range models {
		DB.AutoMigrate(model)
	}
}
