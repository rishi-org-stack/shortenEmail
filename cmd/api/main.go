package main

import (
	"fmt"
	"net/http"
	"shortenEmail/internal/app/auth"
	authdb "shortenEmail/internal/app/auth/db"
	"shortenEmail/internal/app/auth/router"
	"shortenEmail/internal/services/mail"
	authJwt "shortenEmail/internal/util/auth"
	"shortenEmail/internal/util/cache"
	"shortenEmail/internal/util/config"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := "postgresql://pgadmin:password@localhost/shemail?"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	env := config.Init()
	jwtService, err := authJwt.Init(env)
	if err != nil {
		fmt.Println("err->jwtService: ", err)
	}
	authDb := authdb.New()
	mailService, err := mail.NewGmailClient("https://oauth2.googleapis.com/token")
	if err != nil {
		fmt.Println(err)
	}
	rdb, err := cache.Connect()
	if err != nil {
		fmt.Println("err->redis: ", err)
	}
	authService := auth.Init(authDb, db, jwtService, rdb, mailService)
	router.Route(authService)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("server started..")
}
