package mail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"shortenEmail/internal/services"
	"shortenEmail/internal/util/client"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func (gc *gmailClient) GetRedirectUrl() (string, error) {
	// ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return "", errors.New("GetRedirectUrl: " + err.Error())

	}

	config, err := google.ConfigFromJSON(
		b,
		gmail.GmailComposeScope,
		gmail.GmailAddonsCurrentMessageMetadataScope,
		gmail.GmailInsertScope,
		gmail.GmailLabelsScope,
		gmail.GmailReadonlyScope,
		gmail.GmailSendScope,
	)

	if err != nil {
		return "", errors.New("GetRedirectUrl: " + err.Error())
	}

	config.RedirectURL = redirectUrl
	googleRedirectUrl := config.AuthCodeURL(
		"state-token",
		oauth2.AccessTypeOffline,
	)
	return googleRedirectUrl, nil
}

func (gc *gmailClient) GetToken(code,
	grant_code string,
	response chan *services.GetTokenResponse) { // *services.GetTokenResponse,

	response <- nil

	hc := client.NewHttpClient(gc.Url, nil)
	getTokenRequest := &services.GetTokenRequest{
		Code:         code,
		RedirectUri:  redirectUrl,
		ClientID:     gc.clientID,
		ClientSecret: gc.clientSecret,
		GrantType:    grant_code,
	}

	marshledData, err := json.Marshal(getTokenRequest)

	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	res, err := hc.Post(bytes.NewBuffer(marshledData))

	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	returnResponse, err := ioutil.ReadAll(res)
	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	tokenResponse := &services.GetTokenResponse{}

	err = json.Unmarshal(returnResponse, tokenResponse)

	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	response <- tokenResponse

	// return nil
}

func (gc *gmailClient) GetTokenFromRefreshToken(
	refreshToken string,
	response chan *services.RefreshTokenResponse) {

	response <- nil

	hc := client.NewHttpClient(gc.Url, nil)
	getTokenRequest := &services.RefreshTokenRequest{
		ClientID:     gc.clientID,
		RefreshToken: refreshToken,
		ClientSecret: gc.clientSecret,
		GrantType:    refreshGrantCode,
	}

	marshaledData, err := json.Marshal(getTokenRequest)

	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	res, err := hc.Post(bytes.NewBuffer(marshaledData))

	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	returnResponse, err := ioutil.ReadAll(res)
	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	tokenResponse := &services.RefreshTokenResponse{}

	err = json.Unmarshal(returnResponse, tokenResponse)

	if err != nil {
		fmt.Println(errors.New("GetToken: " + err.Error()))
	}

	response <- tokenResponse

	// return nil
}
