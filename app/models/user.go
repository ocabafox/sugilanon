package models

import (
	"strconv"
	"time"

	"github.com/XanderDwyl/sugilanon/app/config"
	jwt "github.com/dgrijalva/jwt-go"
)

// User ....
type User struct {
	ID                int64      `gorm:"AUTO_INCREMENT" json:"id"`
	FacebookId        string     `gorm:"type:varchar(32);unique_index" json:facebook_id"`
	Username          string     `gorm:"type:varchar(32);unique_index" json:"username"`
	Name              string     `gorm:"type:varchar(32)" json:"name"`
	Email             string     `gorm:"type:varchar(130);unique_index" json:"email"`
	VerificationToken string     `json:"verification_token,omitempty"`
	IsVerified        bool       `json:"is_verified,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
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

func (u *User) Create() (User, error) {
	u.Username = "anonymouse" + strconv.Itoa(int(time.Now().UnixNano()))
	u.VerificationToken = "TOKEN"
	u.IsVerified = false
	err := db.Debug().Model(&u).Create(&u).Error

	return *u, err
}

func (u *User) GetUser() ([]User, error) {
	var user []User
	err := db.Debug().Model(&User{}).Where("facebook_id=?", u.FacebookId).Scan(&user).Error

	return user, err
}
