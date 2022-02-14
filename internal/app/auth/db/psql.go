package db

import (
	"fmt"
	"shortenEmail/internal/app/auth"

	"gorm.io/gorm"
)

type authDb struct{}

func New() auth.DB {
	return &authDb{}
}

// func (atb *authDb) FindOrInsert(db *gorm.DB, atr *auth.AuthRequest) (*auth.Account, error) {

// 	tx := db.Where(&auth.Account{Email: atr.Email}).FirstOrCreate(atr)
// 	return atr, tx.Error
// }

func (atb *authDb) Find(db *gorm.DB, atr *auth.AuthRequest) (*auth.Account, error) {
	//
	acc := &auth.Account{}
	tx := db.Where(&auth.Account{Email: atr.Email}).First(acc)

	fmt.Println("look: ", tx.Error != nil)
	return acc, tx.Error
}

func (atb *authDb) Create(db *gorm.DB, atr *auth.Account) (*auth.Account, error) {
	//
	tx := db.Create(atr)
	return atr, tx.Error
}
func (atb *authDb) Update(db *gorm.DB, atr *auth.Account) (*auth.Account, error) {

	tx := db.Where("email=?", atr.Email).Updates(atr)
	return atr, tx.Error
}

func (atb *authDb) Get(db *gorm.DB, id int64) (*auth.Account, error) {

	var authReq = &auth.Account{
		ID: id,
	}
	tx := db.First(authReq)
	return authReq, tx.Error
}
