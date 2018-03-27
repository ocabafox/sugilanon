package models

import (
	"regexp"
	"time"

	"github.com/XanderDwyl/sugilanon/app/config"
	jwt "github.com/dgrijalva/jwt-go"
)

var emailRegex *regexp.Regexp

func init() {
	emailRegex, _ = regexp.Compile(`^[^@]+@[^@]+$`)
}

// User ....
type User struct {
	ID                int64      `gorm:"AUTO_INCREMENT" json:"id"`
	Name              string     `gorm:"type:varchar(32)" json:"name"`
	Email             string     `gorm:"type:varchar(130);unique_index" json:"email,omitempty"`
	VerificationToken string     `json:"verification_token,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}

// UserTable ...
type UserTable struct{}

// GetUserByEmail ...
func (repo *UserTable) GetUserByEmail(email string) (User, error) {
	var user User

	err := db.Debug().Where("email = ?", email).Limit(1).First(&user).Error

	return user, err
}

// CreateJWToken ...
func (u *User) CreateJWToken() (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JWTUser{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})

	tokenString, err := token.SignedString([]byte(config.GetJWTSalt()))

	return tokenString, err
}
