package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"shortenEmail/internal/job"
	"shortenEmail/internal/job/account"
	accDB "shortenEmail/internal/job/account/db"
	"shortenEmail/internal/services/mail"
)

func main() {
	godotenv.Load()
	dsn := "postgresql://pgadmin:password@localhost/shemail"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	accountService := account.New(accDB.New(), db)

	mailService, err := mail.NewGmailClient("https://oauth2.googleapis.com/token")
	if err != nil {
		fmt.Println(err)
	}

	j, _ := job.NewFlusher(context.Background(), accountService, mailService)
	fmt.Println("j ", j)
	j.Run(j.Context)
}
