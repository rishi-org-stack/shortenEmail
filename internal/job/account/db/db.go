package db

import (
	"fmt"
	"gorm.io/gorm"
	"shortenEmail/internal/job/account"
)

type accountDB struct{}

func New() account.DB {
	return &accountDB{}
}

func (adb accountDB) Find(gdb *gorm.DB, expiresIn string) (*[]account.Account, error) {
	fmt.Println(expiresIn)
	acc := &[]account.Account{}
	tx := gdb.Where("expires_on=? or expires_on<?", expiresIn, expiresIn+"11").Find(acc)
	return acc, tx.Error
}

func (adb accountDB) Update(gdb *gorm.DB, acc *account.Account) (*account.Account, error) {
	tx := gdb.Where("email=?", acc.Email).Updates(acc)
	return acc, tx.Error
}
