package main

import (
	"fmt"
	"shortenEmail/internal/app/auth"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://pgadmin:password@localhost/shemail?"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	// tx := db.Exec("drop database shemail")
	tx := db.Exec("drop schema if exists  auth cascade")
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
	}
	tx = db.Exec("create schema if not exists auth")

	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
	}
	err = db.AutoMigrate(&auth.Account{})

	if err != nil {
		fmt.Println(err)
	}
}
