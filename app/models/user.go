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
	FacebookId        string     `gorm:"type:varchar(32)" json:facebook_id"`
	Name              string     `gorm:"type:varchar(32)" json:"name"`
	Username          string     `gorm:"type:varchar(32)" json:"username"`
	Email             string     `gorm:"type:varchar(130)" json:"email"`
	Link              string     `gorm:"type:varchar(130)" json:"link"`
	Gender            string     `gorm:"type:varchar(130)" json:"gender,omitempty"`
	Website           string     `gorm:"type:varchar(130)" json:"website,omitempty"`
	Updated           string     `gorm:"type:varchar(130)" json:"updated"`
	VerificationToken string     `json:"verification_token,omitempty"`
	IsVerified        bool       `json:"is_verified,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at",omitempty"`
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

func (u *User) Update() (User, error) {
	err := db.Debug().Model(&u).Omit("facebook_id", "created_at").Updates(&u).Error

	return *u, err
}

func (u *User) Delete() error {
	err := db.Debug().Model(&u).Where("facebook_id=?", u.FacebookId).Delete(&u).Error

	return err
}

func (u *User) GetUser() (User, error) {
	var user User
	err := db.Debug().Model(&User{}).Where("facebook_id=?", u.FacebookId).Scan(&user).Error

	return user, err
}

func GetUserById(facebookId string) (User, error) {
	var user User
	err := db.Debug().Model(&User{}).Where("facebook_id=?", facebookId).Scan(&user).Error

	return user, err
}
