package account

import (
	"gorm.io/gorm"
	"shortenEmail/internal/services"
	"time"
)

type (
	DB interface {
		Find(db *gorm.DB, expiresIn string) (*[]Account, error)
		Update(db *gorm.DB, atr *Account) (*Account, error)
	}

	Service interface {
		Find(expiresIn string) (*[]Account, error)
		Update(atr *Account) (*Account, error)
	}
	MailService interface {
		GetTokenFromRefreshToken(
			refreshToken string,
			response chan *services.RefreshTokenResponse)
	}

	RedisCache interface {
		Get(string) (string, error)
		Set(string, string, time.Duration) error
	}

	Account struct {
		ID           int64 `json:"id" gorm:"primaryKey"`
		Email        string
		Expired      bool
		ExpiresOn    string
		RefreshToken string
		Status       int
	}
)

func (Account) TableName() string {
	return "auth.account"
}
