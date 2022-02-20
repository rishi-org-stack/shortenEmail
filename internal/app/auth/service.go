package auth

import (
	"context"
	"shortenEmail/internal/services"
	"time"

	"gorm.io/gorm"
)

type (
	DB interface {
		Find(db *gorm.DB, atr *AuthRequest) (*Account, error)
		Create(db *gorm.DB, atr *Account) (*Account, error)
		Update(db *gorm.DB, atr *Account) (*Account, error)
		Get(db *gorm.DB, id int64) (*Account, error)
	}
	Service interface {
		HandleAuth(ctx context.Context, ar *AuthRequest)
		HandleGetCode(ctx context.Context, email string)
		HandleCode(code, email string)
	}

	MailService interface {
		GetRedirectUrl() (string, error)
		GetToken(code, grant_code string, response chan *services.GetTokenResponse)
	}
	redisCache interface {
		Get(string) (string, error)
		Set(string, string, time.Duration) error
	}

	tokenGenrator interface {
		GenrateToken(id int, accessToken string) (string, error)
	}
	AuthRequest struct {
		Email string `json:"email" gorm:"not null"`
	}

	Account struct {
		ID           int64 `json:"id" gorm:"primaryKey"`
		Email        string
		Expired      bool
		ExpiresOn    string
		RefreshToken string
		Status       int
	}

	TokenGenratorInterface interface {
		GenrateToken(int, string, string) (string, error)
	}
)

const (
	google      = "http://localhost:8080/auth/google"
	grant_type  = "authorization_code"
	confirmed   = 1
	unconfirmed = 0
)

func (Account) TableName() string {
	return "auth.account"
}
