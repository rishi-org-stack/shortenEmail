package job

import (
	"context"
	"fmt"
	"log"
	"shortenEmail/internal/job/account"
	"strconv"
	"time"

	"go.uber.org/atomic"
)

type Flusher struct {
	context.Context
	finish         func()
	accountService account.Service
	cache          account.RedisCache
	enabled        *atomic.Bool
	mailService    account.MailService
	// db *sqlx.DB
}

func NewFlusher(
	ctx context.Context,
	service account.Service,
	mailService account.MailService) (*Flusher, error) {
	job := &Flusher{
		accountService: service,
		mailService:    mailService,
		enabled:        atomic.NewBool(true),
	}
	job.Context, job.finish = context.WithCancel(ctx)
	return job, nil
}

func (flusher *Flusher) Run(ctx context.Context) {
	log.Println("Started background flusher")

	//complete := make(chan string)
	//for {
	//	time.Sleep(time.Second * 2)
	//}
	//for {
	//	x, open := <-complete
	//	if !open {
	//		fmt.Println("oops")
	//	}
	//	fmt.Println("x is", x)
	//
	//}

	//go func() {
	//	for {
	//		x, open := <-complete
	//		if !open {
	//			fmt.Println("oops")
	//		}
	//		fmt.Println("x is", x)
	//
	//	}
	//}()
	for {
		flusher.flush()
		time.Sleep(10 * time.Second)

	}
	//flusher.flush()

}

func (flusher *Flusher) Finish(c chan string) {
	close(c)
}
func (flusher *Flusher) flush() {
	accounts, err := flusher.accountService.Find(strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		fmt.Println("no")
	}
	fmt.Println(accounts)
	//const idPrefix = "USER_"
	//responseChan := make(chan *services.RefreshTokenResponse)
	//
	//for _, acc := range *accounts {
	//	go flusher.mailService.GetTokenFromRefreshToken(acc.RefreshToken, responseChan)
	//}
	//for i := 0; i < len(*accounts); i++ {
	//	response, open := <-responseChan
	//	if !open {
	//		fmt.Println("no")
	//	}
	//	fmt.Println("ok")
	//	if response != nil {
	//
	//		fmt.Println(response)
	//err := flusher.cache.Set(idPrefix+acc.Email, response.AccessToken,
	//	time.Duration(time.Second*time.Duration(response.ExpiresIn)))
	//if err != nil {
	//	fmt.Println("no")
	//
	//}
	//res, err := flusher.cache.Get(idPrefix + acc.Email)
	//if err != nil {
	//	fmt.Println("no")
	//
	//}
	//fmt.Println(res)

	//}

	//}
	//result, err := flusher.accountService.Update(&account.Account{
	//	Email:     "ok@gmail.com",
	//	ExpiresOn: time.Now().Unix(),
	//})
	//if err != nil {
	//	c <- "no"
	//}
	//fmt.Println("result.expires_on ", result.ExpiresOn)
	//c <- "ok"
}
