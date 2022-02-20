package account

import (
	"fmt"
	"gorm.io/gorm"
)

type authService struct {
	authDB DB
	gdb    *gorm.DB
}

func New(db DB, gdb *gorm.DB) Service {
	return &authService{
		db,
		gdb,
	}

}

func (as authService) Find(expiresIn string) (*[]Account, error) {
	acc, err := as.authDB.Find(as.gdb, expiresIn)
	if err != nil {
		return nil, fmt.Errorf("find: %w", err)
	}
	return acc, nil
}
func (as authService) Update(acc *Account) (*Account, error) {
	acc, err := as.authDB.Update(as.gdb, acc)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}
	return acc, nil
}
