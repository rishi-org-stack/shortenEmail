package auth

import (
	"context"
	"fmt"
	"net/http"
	"shortenEmail/internal/services"
	"shortenEmail/internal/util"
	"shortenEmail/internal/util/cache"
	"time"

	"gorm.io/gorm"
)

//error handle properly
const (
	source                  = "AUTH"
	AUTH_INSERT_ERROR       = source + "_INSERT_ERROR"
	AUTH_SERVER_ERROR       = source + "_SERVER_ERROR"
	AUTH_BAD_REQUEST        = source + "_BAD_REQUEST"
	AUTH_OTP_INSERT_ERROR   = source + "_OTP_INSERT_ERROR"
	AUTH_UNAUTHORIZED_ERROR = source + "_INSERT_ERROR"
)

type authService struct {
	gdb         *gorm.DB
	authData    DB
	mailService MailService
	// jwtSer      TokenGenratorInterface
	// config      *config.Env
}

var OTP string

func Init(db DB, gdb *gorm.DB, mailservice MailService) Service {
	return &authService{
		gdb:      gdb,
		authData: db,
		// jwtSer:      js,
		// config:      config,
		mailService: mailservice,
	}
}

func (authServ authService) HandleAuth(ctx context.Context, ar *AuthRequest) { // (*apiRes.Response, apiError.ApiErrorInterface) {

	res, err := authServ.authData.Find(authServ.gdb, ar)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("look")
		fmt.Println(ar)
		req := (ctx.Value("env")).(map[interface{}]interface{})["r"].(*http.Request)
		w := (ctx.Value("env")).(map[interface{}]interface{})["w"].(http.ResponseWriter)

		http.Redirect(w, req, "http://localhost:8080/auth/getCode?email="+ar.Email, http.StatusTemporaryRedirect)
		// 	return
	}

	rdb, err := cache.Connect()

	if err != nil {
		fmt.Println(err)
	}

	accessToken, err := rdb.Get("USER_" + res.Email)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ttoken", accessToken)

}

func (authServ authService) HandleGetCode(ctx context.Context, email string) {
	req := (ctx.Value("env")).(map[interface{}]interface{})["r"].(*http.Request)
	w := (ctx.Value("env")).(map[interface{}]interface{})["w"].(http.ResponseWriter)

	acc := &Account{
		Email:  email,
		Status: unconfirmed,
	}

	_, err := authServ.authData.Create(authServ.gdb, acc)

	if err != nil {
		fmt.Println(err)
	}

	url, err := authServ.mailService.GetRedirectUrl()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done")
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (authServ authService) HandleCode(code, email string) {
	if email == "" {
		fmt.Println("no mail")
		return
	}
	response := make(chan *services.GetTokenResponse)

	go authServ.mailService.GetToken(code, grant_type, response)

	for {
		getTokenRes, open := <-response
		if !open {
			fmt.Println("not open")
			return
		}
		fmt.Println("xxxxxx")
		if getTokenRes != nil {

			fmt.Println(getTokenRes)
			acc := &Account{
				Email:        email,
				Expired:      false,
				RefreshToken: getTokenRes.RefreshToken,
				ExpiresOn:    util.TimeToEpoch(int64(getTokenRes.ExpiresIn)),
				Status:       confirmed,
			}

			acc, err := authServ.authData.Update(authServ.gdb, acc)
			if err != nil {
				fmt.Println(err)
			}
			rdb, err := cache.Connect()
			if err != nil {
				fmt.Println(err)
			}
			err = rdb.Set(
				"USER_"+email, getTokenRes.AccessToken,
				time.Duration(time.Second*time.Duration(getTokenRes.ExpiresIn)))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(acc)

			return
		}
	}

}
