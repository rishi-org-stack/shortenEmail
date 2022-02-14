package main

import (
	"context"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"log"
	"net/http"
	"os"

	// "strconv"
	"shortenEmail/internal/app/auth"
	authdb "shortenEmail/internal/app/auth/db"
	"shortenEmail/internal/app/auth/router"
	"shortenEmail/internal/services/mail"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/gmail/v1"
	// "google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	godotenv.Load()
	dsn := "postgresql://pgadmin:password@localhost/shemail?"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	authDb := authdb.New()
	mailService, err := mail.NewgmailClient("https://oauth2.googleapis.com/token")
	if err != nil {
		fmt.Println(err)
	}

	authService := auth.Init(authDb, db, mailService)
	router.Route(authService)
	http.ListenAndServe(":8080", nil)

}

// package main

// import (
// 	// "shortenEmail/internal/app/auth/router"
// 	"net/http"
// 	// "net/mail"
// 	"shortenEmail/internal/app/auth/router"
// 	"shortenEmail/internal/app/mail"
// )

// func main() {
// 	router.Route()
// 	mail.Route()
// 	http.ListenAndServe(":8080", nil)
// }
